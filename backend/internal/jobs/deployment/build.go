package deploymentjob

import (
	"context"
	"database/sql"
	"fmt"
	"os/exec"
	"path"

	"github.com/Roshan-anand/godploy/internal/db"
	deploymentqueue "github.com/Roshan-anand/godploy/internal/jobs/deployment/queue"
	logbrokerqueue "github.com/Roshan-anand/godploy/internal/jobs/logbroker/queue"
	"github.com/Roshan-anand/godploy/internal/lib/types"
)

func getDockerBuildCmd(d *deploymentqueue.BuildJobData) *exec.Cmd {
	// 	"--secret", "id=npm_token,src=/tmp/npm_token",
	// 	"--secret", "id=github_token,src=/tmp/github_token",

	cmd := exec.Command("docker", "buildx", "build")

	if d.DockerFilePath != "" {
		cmd.Args = append(cmd.Args, "--file", d.DockerFilePath)
	}

	for _, arg := range d.BuildArgs {
		cmd.Args = append(cmd.Args, "--build-arg", arg)
	}

	// TODO : add build secrets to the cmd

	if d.ImgName != "" {
		cmd.Args = append(cmd.Args, "--tag", d.ImgName)
	}

	if d.DockerBuildStage != "" {
		cmd.Args = append(cmd.Args, "--target", d.DockerBuildStage)
	}

	// create a tar achive of the code folder
	dockerCtxPath := path.Join(d.StorePath + d.DockerContextPath)
	cmd.Args = append(cmd.Args, dockerCtxPath)

	return cmd
}

// responsible for pulling code and storing it local
func (w *worker) BuildWorker(ctx context.Context, data chan *deploymentqueue.BuildJobData) {
	for {
		select {
		case d, ok := <-data:
			if !ok {
				fmt.Println("BuildWorker: data channel closed, exiting")
				return
			}

			fmt.Println("BuildWorker: started working ...")
			l := w.Server.LogBrokerQ

			// generate a new docker build cmd
			buildCmd := getDockerBuildCmd(d)

			if err := runWorkerCmd(l, d.DeploymentID, buildCmd); err != nil {
				fmt.Printf("PullWorker: error running command: %v\n", err)
				l.EndLogs(&logbrokerqueue.EndLogData{
					DeploymentID: d.DeploymentID,
					Status:       types.DeploymentError,
				})
				continue
			}

			// update the deployment with the built image name
			if err := w.Server.DB.Queries.SetDeploymentImageName(w.qCtx, db.SetDeploymentImageNameParams{
				ID: d.DeploymentID,
				ImageName: sql.NullString{
					Valid:  true,
					String: d.ImgName,
				},
			}); err != nil {
				fmt.Printf("BuildWorker: error updating deployment image name: %v\n", err)
				l.EndLogs(&logbrokerqueue.EndLogData{
					DeploymentID: d.DeploymentID,
					Status:       types.DeploymentError,
				})
				continue
			}

			// set a deploy worker
			w.Server.DeploymentQ.EnqueueDeployJob(&deploymentqueue.DeployJobData{
				DeploymentID:     d.DeploymentID,
				SwarmServiceName: d.SwarmServiceName,
				ImgName:          d.ImgName,
				Env:              d.Env,
			})

		case <-ctx.Done():
			fmt.Println("BuildWorker: context cancelled, exiting")
			return
		}
	}
}
