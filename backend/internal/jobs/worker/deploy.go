package worker

import (
	"context"
	"fmt"

	"github.com/Roshan-anand/godploy/internal/jobs/queue"
)

// responsible for pulling code and storing it local
func (w *worker) DeployWorker(ctx context.Context, data chan *queue.DeployJobData) {
	fmt.Println("DeployWorker: started")
	for {
		select {
		case d, ok := <-data:
			if !ok {
				fmt.Println("DeployWorker: data channel closed, exiting")
				return
			}

			fmt.Printf("deployWorker: received job data: %+v\n", d)

		case <-ctx.Done():
			fmt.Println("DeployWorker: context cancelled, exiting")
			return
		}
	}
}
