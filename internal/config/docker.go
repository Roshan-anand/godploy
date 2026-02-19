package config

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/docker/docker/client"
)

func InitDockerClient() (*client.Client, error) {
	c, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		return nil, err
	}

	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	p, err := c.Ping(ctx)
	if err != nil {
		if closeErr := c.Close(); closeErr != nil {
			return nil, errors.Join(err, closeErr)
		}
		return nil, err
	}

	fmt.Println("connected docker :", p.APIVersion)

	return c, nil
}

func (s *Server) CloseDockerClient() error {
	fmt.Println("closing docker client connection")
	return s.Docker.Close()
}
