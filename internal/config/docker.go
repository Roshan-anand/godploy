package config

import (
	"context"
	"errors"
	"fmt"
	"time"

	"github.com/moby/moby/client"
)

func InitDockerClient() (*client.Client, error) {
	c, err := client.New(client.FromEnv)
	if err != nil {
		return nil, err
	}

	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	p, err := c.Ping(ctx, client.PingOptions{})
	if err != nil {
		if closeErr := c.Close(); closeErr != nil {
			return nil, errors.Join(err, closeErr)
		}
		return nil, err
	}

	fmt.Println("connected docker :", p.APIVersion)

	// initialize swarm mode if not initialized
	if _, err = c.SwarmInspect(context.Background(), client.SwarmInspectOptions{}); err != nil {
		if _, err := c.SwarmInit(context.Background(), client.SwarmInitOptions{
			AdvertiseAddr: "127.0.0.1",
			ListenAddr:    "0.0.0.0",
		}); err != nil {
			return nil, fmt.Errorf("failed to initialize swarm mode : %w", err)
		}
	}

	return c, nil
}

func (s *Server) CloseDockerClient() error {
	fmt.Println("closing docker client connection")
	return s.Docker.Close()
}
