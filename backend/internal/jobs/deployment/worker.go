package deployjob

import (
	"context"
	"database/sql"
	"fmt"
	"os"
	"os/exec"
	"path"

	"github.com/Roshan-anand/godploy/internal/db"
	"github.com/Roshan-anand/godploy/internal/jobs/logbroker"
	"github.com/Roshan-anand/godploy/internal/lib/types"
	"github.com/docker/docker/api/types/swarm"
	"github.com/google/uuid"
)

func (d *DeploymentService) runDeploymentPipeline(ctx context.Context, data *DeploymentServiceParams) {

	d.log.PublishLog(&logbroker.PubData{
		ID:  data.DeploymentID,
		Msg: getTitle("Pulling code from" + data.Url),
	})

	// update the deployment status to building
	if err := d.queries.UpdateDeploymentStatus(d.qCtx, db.UpdateDeploymentStatusParams{
		Status: types.DeploymentBuilding,
		ID:     data.DeploymentID,
	}); err != nil {
		fmt.Printf("PullWorker: error updating deployment status: %v\n", err)
	}

	outputPath := path.Join(d.codeStoreDir, data.SwarmService)
	repoUrl := fmt.Sprintf("https://oauth2:%s@%s", data.Token, data.Url)
	cmd := exec.Command("git", "clone", "--branch", data.Branch, "--depth", "1", repoUrl, "outputPath")

	if err := runWorkerCmd(d.log, data.DeploymentID, cmd, "pull"); err != nil {
		fmt.Printf("PullWorker: error running command: %v\n", err)
		d.log.EndLogs(&logbroker.EndLogData{
			DeploymentID: data.DeploymentID,
			Status:       types.DeploymentError,
		})
		return // TODO : trigger retry logic
	}

	fmt.Println("finished pulling :", data.Url)

	d.log.PublishLog(&logbroker.PubData{
		ID:  data.DeploymentID,
		Msg: getTitle("Building the image " + data.ImgName),
	})

	// generate a new docker build cmd
	buildCmd := getDockerBuildCmd(&dockerBuildCmdData{
		dockerFilePath:    data.DockerFilePath,
		dockerContextPath: data.DockerContextPath,
		dockerBuildStage:  data.DockerBuildStage,
		imgName:           data.ImgName,
		buildArgs:         data.BuildArgs,
		outputPath:        outputPath,
	})

	if err := runWorkerCmd(d.log, data.DeploymentID, buildCmd, "build"); err != nil {
		fmt.Printf("BuildWorker: error running command: %v\n", err)
		d.log.EndLogs(&logbroker.EndLogData{
			DeploymentID: data.DeploymentID,
			Status:       types.DeploymentError,
		})
		return // TODO : trigger retry logic
	}

	// update the deployment with the built image name
	if err := d.queries.SetDeploymentImageName(d.qCtx, db.SetDeploymentImageNameParams{
		ID: data.DeploymentID,
		Image: sql.NullString{
			Valid:  true,
			String: data.ImgName,
		},
	}); err != nil {
		fmt.Printf("BuildWorker: error updating deployment image name: %v\n", err)
		d.log.EndLogs(&logbroker.EndLogData{
			DeploymentID: data.DeploymentID,
			Status:       types.DeploymentError,
		})
		return // TODO : trigger retry logic
	}

	// remove the code folder
	go os.RemoveAll("outputPath")

	fmt.Println("finished building :", data.ImgName)

	switch data.JobType {
	case DeployJob:
		// set a deploy worker
		// w.Server.DeploymentQ.EnqueueDeployJob(&deploymentqueue.DeployJobData{
		// 	DeploymentID: d.DeploymentID,
		// 	SwarmService: d.SwarmService,
		// 	ImgName:      d.ImgName,
		// 	Env:          d.Env,
		// 	IsPublic:     d.IsPublic,
		// 	NetworkName:  d.NetworkName,
		// })

	case ReDeployJob:
		// w.Server.DeploymentQ.EnqueueRedeployJob(&deploymentqueue.RedeployJobData{
		// 	DeploymentID: d.DeploymentID,
		// 	SwarmService: d.SwarmService,
		// 	ImgName:      d.ImgName,
		// 	Env:          d.Env,
		// 	IsPublic:     d.IsPublic,
		// 	NetworkName:  d.NetworkName,
		// })
	default:
		fmt.Printf("BuildWorker: unknown job type: %v\n", data.JobType)
	}
}

type deployData struct {
	deploymentID uuid.UUID
	swarmService string
	networkName  string
	isPublic     bool
	env          []string
	imgName      string
}

func (d *DeploymentService) deploy(data *deployData) {

	d.log.PublishLog(&logbroker.PubData{
		ID:  data.deploymentID,
		Msg: getTitle("Deploying  the service " + data.swarmService),
	})

	// create network if not exist
	if err := d.docker.CreateNetwork(data.networkName); err != nil {
		fmt.Printf("DeployWorker: error creating network: %v\n", err)
		d.log.EndLogs(&logbroker.EndLogData{
			DeploymentID: data.deploymentID,
			Status:       types.DeploymentError,
			Message:      err.Error(),
		})

		return // TODO : trigger retry logic
	}

	// get the service spec
	spec := getBaseSpec(data.imgName, data.networkName, data.swarmService, data.env, data.isPublic)

	_, err := d.docker.Client.ServiceCreate(context.Background(), *spec, swarm.ServiceCreateOptions{})
	if err != nil {
		fmt.Printf("DeployWorker: error creating service: %v\n", err)
		d.log.EndLogs(&logbroker.EndLogData{
			DeploymentID: data.deploymentID,
			Status:       types.DeploymentError,
			Message:      err.Error(),
		})

		return // TODO : trigger retry logic
	}

	fmt.Println("finished deploying :", data.swarmService)
	d.log.EndLogs(&logbroker.EndLogData{
		DeploymentID: data.deploymentID,
		Status:       types.DeploymentReady,
		Message:      getTitle("successfully deployed"),
	})
}

func (d *DeploymentService) redeploy(data *deployData) {
	d.log.PublishLog(&logbroker.PubData{
		ID:  data.deploymentID,
		Msg: "Redeploying  the service " + data.swarmService,
	})

	// get the swarm service spec
	res, _, err := d.docker.Client.ServiceInspectWithRaw(context.Background(), data.swarmService, swarm.ServiceInspectOptions{})
	if err != nil {
		fmt.Printf("DeployWorker: error inspecting service: %v\n", err)
		d.log.EndLogs(&logbroker.EndLogData{
			DeploymentID: data.deploymentID,
			Status:       types.DeploymentError,
			Message:      err.Error(),
		})

		return // TODO : trigger retry logic
	}
	version := res.Version
	spec := res.Spec

	// update the image and env
	spec.TaskTemplate.ContainerSpec.Image = data.imgName
	if len(data.env) > 0 {
		spec.TaskTemplate.ContainerSpec.Env = data.env
	}

	// update the service with the new spec
	if _, err := d.docker.Client.ServiceUpdate(context.Background(), data.swarmService, version, spec, swarm.ServiceUpdateOptions{}); err != nil {
		fmt.Printf("DeployWorker: error updating service: %v\n", err)
		d.log.EndLogs(&logbroker.EndLogData{
			DeploymentID: data.deploymentID,
			Status:       types.DeploymentError,
			Message:      err.Error(),
		})

		return // TODO : trigger retry logic
	}

	// end the logs
	d.log.EndLogs(&logbroker.EndLogData{
		DeploymentID: data.deploymentID,
		Status:       types.DeploymentReady,
		Message:      getTitle("successfully redeployed"),
	})
}
