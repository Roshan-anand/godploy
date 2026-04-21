package worker

import (
	"context"
	"fmt"

	"github.com/Roshan-anand/godploy/internal/jobs/queue"
)

// responsible for pulling code and storing it local
func (w *worker) PullWorker(ctx context.Context, data chan *queue.PullJobData) {
	fmt.Println("PullWorker: started")
	for {
		select {
		case d, ok := <-data:
			if !ok {
				fmt.Println("PullWorker: data channel closed, exiting")
				return
			}

			fmt.Printf("PullWorker: received job data: %+v\n", d)

			// repoURL := fmt.Sprintf("https://oauth2:%s@github.com/%s/%s.git", pData.Token, pData.Owner, pData.Repo)

			// outputPath := fmt.Sprintf("/etc/godploy/code/%s-%s-%s", pData.Owner, pData.Repo, pData.Branch)

			// _ = exec.Command("git", "clone", "--branch", pData.Branch, "--depth", "1", repoURL, outputPath)
		case <-ctx.Done():
			fmt.Println("PullWorker: context cancelled, exiting")
			return
		}
	}
}
