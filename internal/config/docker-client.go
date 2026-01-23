package config

import (
	"context"
	"fmt"
	"time"

	"github.com/docker/docker/client"
)

func getDockerClient() (*client.Client, error) {
	client, err := client.NewClientWithOpts(client.FromEnv, client.WithAPIVersionNegotiation())
	if err != nil {
		// TODO: handle error properly by
		// - logging the error let user decide what to do.
		// - keep docker client as optional and fallback to other providers.
		return nil, err
	}

	ctx, cancle := context.WithTimeout(context.Background(), time.Second*5)
	defer cancle()

	// ping the docker server
	_, err = client.Ping(ctx)
	if err != nil {
		return nil, err
	}

	info, err := client.Info(ctx)
	if err != nil {
		return nil, err
	}
	
	fmt.Printf("server connected to docker at %s \n os: %s\n", info.DockerRootDir, info.OperatingSystem, )
	return client, nil
}
