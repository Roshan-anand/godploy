package worker

import "github.com/Roshan-anand/godploy/internal/config"

type worker struct {
	Server *config.Server
}

func NewWorker(s *config.Server) *worker {
	return &worker{
		Server: s,
	}
}
