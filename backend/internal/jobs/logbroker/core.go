package logbroker

import (
	"context"
	"fmt"

	"github.com/Roshan-anand/godploy/internal/config"
	logbrokerqueue "github.com/Roshan-anand/godploy/internal/jobs/logbroker/queue"
)

type LogsBroker struct {
	Server *config.Server
}

func InitLogsBroker(s *config.Server) *LogsBroker {
	return &LogsBroker{
		Server: s,
	}
}

func (job *LogsBroker) LogsBrokerJob(ctx context.Context, pub chan *logbrokerqueue.PubData) {
	fmt.Println("LogBroker: started")
	for {
		select {
		case p, ok := <-pub:
			if !ok {
				fmt.Println("Publisher channel closed, exiting logBroker")
				return
			}
			println("log broker recived message : ", p.Msg)

			// check for subscribers
			for _, sub := range job.Server.LogBrokerQ.Subscribers {
				if sub.DeploymentID == p.ID {
					sub.SSE.SendSSE("event", "data")
				}
			}

			// TODO : send it to buffer

		case <-ctx.Done():
			fmt.Println("Context cancelled, exiting logBroker")
			return
		}
	}
}
