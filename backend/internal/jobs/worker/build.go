package worker

import (
	"context"
	"fmt"

	"github.com/Roshan-anand/godploy/internal/jobs/queue"
)

// responsible for pulling code and storing it local
func (w *worker) BuildWorker(ctx context.Context, data chan *queue.BuildJobData) {
	fmt.Println("BuildWorker: started")
	for {
		select {
		case d, ok := <-data:
			if !ok {
				fmt.Println("BuildWorker: data channel closed, exiting")
				return
			}

			fmt.Printf("BuildWorker: received job data: %+v\n", d)

		case <-ctx.Done():
			fmt.Println("BuildWorker: context cancelled, exiting")
			return
		}
	}
}
