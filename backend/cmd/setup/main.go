package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/Roshan-anand/godploy/internal/config"
	// "github.com/moby/moby/client"
)

func runCommand(name string, args ...string) error {
	cmd := exec.Command(name, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// to setup local development env
func main() {
	c, err := config.InitDockerClient()
	if err != nil {
		fmt.Println(err)
		return
	}

	host := c.DaemonHost()
	os.Setenv("DOCKER_HOST", host)

	switch os.Args[1] {
	case "setup":
		if err := runCommand("docker", "stack", "deploy", "-c", "../dynamic/compose.yaml", "godploy"); err != nil {
			fmt.Println("failed to setup traefik stack :", err)
			return
		}
	case "dev":
		if err := runCommand("docker", "stack", "deploy", "-c", "../docker/compose.dev.yaml", "godploy"); err != nil {
			fmt.Println("failed to setup godploy stack :", err)
			return
		}
	default:
		fmt.Println("invalid command")
	}
}
