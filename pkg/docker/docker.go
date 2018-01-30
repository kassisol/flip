package docker

import (
	"context"
	"fmt"

	"github.com/docker/docker/api/types"
	"github.com/docker/docker/api/types/filters"
	"github.com/docker/docker/client"
)

type DockerConfig struct {
	Client *client.Client
	NodeID string
}

func NewDockerClient() (*DockerConfig, error) {
	cli, err := client.NewEnvClient()
	if err != nil {
		return &DockerConfig{}, err
	}

	return &DockerConfig{Client: cli}, nil
}

func (c *DockerConfig) GetNodeID() error {
	info, err := c.Client.Info(context.Background())
	if err != nil {
		return err
	}

	c.NodeID = info.Swarm.NodeID

	return nil
}

func (c *DockerConfig) IsServiceRunning() (bool, error) {
	filterID := fmt.Sprintf("{\"label\": {\"flip.enabled=true\": true}, \"node\": {\"%s\": true}, \"desired-state\": {\"running\": true}}", c.NodeID)
	f, err := filters.FromParam(filterID)
	if err != nil {
		return false, err
	}

	opts := types.TaskListOptions{
		Filters: f,
	}

	tasks, err := c.Client.TaskList(context.Background(), opts)
	if err != nil {
		return false, err
	}

	if len(tasks) == 1 {
		return true, nil
	}

	return false, nil
}
