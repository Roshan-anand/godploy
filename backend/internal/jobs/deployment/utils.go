package deploymentjob

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os/exec"

	logbrokerqueue "github.com/Roshan-anand/godploy/internal/jobs/logbroker/queue"
	"github.com/creack/pty"
	"github.com/docker/docker/api/types/swarm"
	"github.com/google/uuid"
)

// returns a formatted title string for the logs
func getTitle(msg string) string {
	return fmt.Sprintf("\n-----------------------------------\n\n %s \n------------------------------------\n", msg)
}

// scans the reader line by line and publish the logs
func scanAndPublish(l *logbrokerqueue.LogBrokerQueue, dID uuid.UUID, r io.Reader) {
	reader := bufio.NewReader(r)

	for {
		line, err := reader.ReadString('\n')
		if len(line) > 0 {
			l.PublishLog(&logbrokerqueue.PubData{
				ID:  dID,
				Msg: line,
			})
		}

		if err != nil {
			if !errors.Is(err, io.EOF) {
				fmt.Println("stdout read error:", err)
			}
			break
		}
	}
}

// runs the given cmd in a psuedo terminal and publishes the logs to the log broker
func runWorkerCmd(l *logbrokerqueue.LogBrokerQueue, dID uuid.UUID, cmd *exec.Cmd, worker string) error {
	ptmx, err := pty.Start(cmd)
	if err != nil {
		return fmt.Errorf("%s:err:pty:start: %v", worker, err)
	}
	defer ptmx.Close()

	go scanAndPublish(l, dID, ptmx)

	if err := cmd.Wait(); err != nil {
		return fmt.Errorf("%s:err:cmd:wait: %v\n", worker, err)
	}
	return nil
}

// returns a base service spec for the given parameters
func getBaseSpec(imgName string, networkName string, swarmName string, env []string, isPublic bool) *swarm.ServiceSpec {

	spec := &swarm.ServiceSpec{
		Annotations: swarm.Annotations{
			Name: swarmName,
			Labels: map[string]string{
				fmt.Sprintf("traefik.http.routers.%s.entrypoints", swarmName):               "websecure",
				fmt.Sprintf("traefik.http.services.%s.loadbalancer.server.port", swarmName): "80",
				fmt.Sprintf("traefik.http.routers.%s.tls.certresolver", swarmName):          "le",
				"traefik.constraint-label": "head-proxy",
			},
		},

		TaskTemplate: swarm.TaskSpec{
			ContainerSpec: &swarm.ContainerSpec{
				Image: imgName,
				TTY:   false,
			},

			RestartPolicy: &swarm.RestartPolicy{
				Condition: swarm.RestartPolicyConditionAny,
			},

			Networks: []swarm.NetworkAttachmentConfig{
				{
					Target: networkName,
				},
			},
		},
	}

	// if the service is public connect to traefik
	if isPublic {
		spec.TaskTemplate.Networks = append(spec.TaskTemplate.Networks, swarm.NetworkAttachmentConfig{
			Target: "godploy_traefik_proxy",
		})
		spec.Annotations.Labels["traefik.enable"] = "true"
	} else {
		spec.Annotations.Labels["traefik.enable"] = "false"
	}

	// if env avalable
	if len(env) > 0 {
		spec.TaskTemplate.ContainerSpec.Env = env
	}

	return spec
}
