package deploymentjob

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"os/exec"

	logbrokerqueue "github.com/Roshan-anand/godploy/internal/jobs/logbroker/queue"
	"github.com/creack/pty"
	"github.com/google/uuid"
)

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
