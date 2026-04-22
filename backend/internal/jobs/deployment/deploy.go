package deploymentjob

import (
	"context"
	"fmt"
	"time"

	deploymentqueue "github.com/Roshan-anand/godploy/internal/jobs/deployment/queue"
	logbrokerqueue "github.com/Roshan-anand/godploy/internal/jobs/logbroker/queue"
	"github.com/Roshan-anand/godploy/internal/lib"
)

// responsible for pulling code and storing it local
func (w *worker) DeployWorker(ctx context.Context, data chan *deploymentqueue.DeployJobData) {
	fmt.Println("DeployWorker: started")
	for {
		select {
		case _, ok := <-data:
			if !ok {
				fmt.Println("DeployWorker: data channel closed, exiting")
				return
			}

			fmt.Println("DeployWorker: started working ...")

			for i := range 5 {
				w.Server.LogBrokerQ.PublishLog(&logbrokerqueue.PubData{
					ID:  lib.NewID(),
					Msg: fmt.Sprintf("deploy : %v", i),
				})
				time.Sleep(1 * time.Second)
			}

			fmt.Printf("DeployWorker: finished working ...")
		case <-ctx.Done():
			fmt.Println("DeployWorker: context cancelled, exiting")
			return
		}
	}
}
