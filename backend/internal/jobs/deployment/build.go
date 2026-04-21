package deploymentjob

import (
	"context"
	"fmt"
	"time"

	deploymentqueue "github.com/Roshan-anand/godploy/internal/jobs/deployment/queue"
	logbrokerqueue "github.com/Roshan-anand/godploy/internal/jobs/pubsub/queue"
	"github.com/Roshan-anand/godploy/internal/lib"
)

// responsible for pulling code and storing it local
func (w *worker) BuildWorker(ctx context.Context, data chan *deploymentqueue.BuildJobData) {
	fmt.Println("BuildWorker: started")
	for {
		select {
		case _, ok := <-data:
			if !ok {
				fmt.Println("BuildWorker: data channel closed, exiting")
				return
			}

			fmt.Println("BuildWorker: started working ...")

			for i := range 5 {
				w.Server.LogBrokerQ.PublishLog(&logbrokerqueue.PubData{
					ID:  lib.NewID(),
					Msg: fmt.Sprintf("build : %v", i),
				})
				time.Sleep(1 * time.Second)
			}

			w.Server.DeploymentQ.EnqueueDeployJob(&deploymentqueue.DeployJobData{})
		case <-ctx.Done():
			fmt.Println("BuildWorker: context cancelled, exiting")
			return
		}
	}
}
