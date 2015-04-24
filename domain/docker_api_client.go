package domain

import (
	"errors"
)

type Container struct {
	Command string   `json:"Command"`
	Created int      `json:"Created"`
	ID      string   `json:"Id"`
	Image   string   `json:"Image"`
	Names   []string `json:"Names"`
	Ports   []struct {
		IP          string `json:"IP"`
		Privateport int    `json:"PrivatePort"`
		Publicport  int    `json:"PublicPort"`
		Type        string `json:"Type"`
	} `json:"Ports"`
	Status string `json:"Status"`
}

type IDockerAPIClientManager interface {
	ListContainers(socketPath string) ([]*Container, error)
}

type DockerAPIClientManager struct {
	InjectedDockerAPIClientManager IDockerAPIClientManager
}

func (manager *DockerAPIClientManager) ListContainers(socketPath string) ([]*Container, error) {
	if manager.InjectedDockerAPIClientManager == nil {
		return nil, errors.New("Injected DockerAPIClientManager cannot be null")
	}
	containers, err := manager.InjectedDockerAPIClientManager.ListContainers(socketPath)

	return containers, err
}
