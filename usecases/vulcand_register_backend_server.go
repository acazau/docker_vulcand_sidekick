package usecases

import (
	"fmt"
	"github.com/acazau/docker_vulcand_sidekick/domain"
	"github.com/acazau/docker_vulcand_sidekick/infrastructure"
	"log"
	"strings"
	"unicode/utf8"
)

var (
	dockerAPIClient  = new(domain.DockerAPIClientManager)
	vulcandAPIClient = new(domain.VulcandAPIClientManager)
)

type Vulcand_Register_Backend_Server_Repository struct {
	useCaseManager UseCaseManager
}

func init() {
	dockerAPIRepo := &infrastructure.DockerAPIClient_HTTP_Repository{}
	vulcandAPIRepo := &infrastructure.VulcandAPIClient_HTTP_Repository{}
	dockerAPIClient.InjectedDockerAPIClientManager = dockerAPIRepo
	vulcandAPIClient.InjectedVulcandAPIClientManager = vulcandAPIRepo
}

func (manager *Vulcand_Register_Backend_Server_Repository) Execute(config *domain.Config) error {

	containers, err := dockerAPIClient.ListContainers(config.DockerAPIEndpoint)
	if err != nil {
		log.Println(fmt.Sprintf("Error getting containers %s", err.Error()))
		return err
	}

	for _, container := range containers {
		log.Println(fmt.Sprintf("Image %s", container.Image))
		for _, port := range container.Ports {
			if port.Type == "tcp" && port.PublicPort > 0 {
				log.Println(fmt.Sprintf("Public Port %d Type %s", port.PublicPort, port.Type))

				backend := new(domain.Backend)
				backend.Backend.ID = strings.Replace(container.Image, "/", "_", -1)
				backend.Backend.Type = "http"

				eBackend, err := vulcandAPIClient.GetBackendById(config.VulcandAPIEndpoint, backend.Backend.ID)
				if err != nil {
					log.Println(fmt.Sprintf("Error Finding backend %s", err.Error()))
				}

				if eBackend == nil || utf8.RuneCountInString(eBackend.Backend.ID) <= 0 {
					eBackend, err = vulcandAPIClient.UpsertBackend(config.VulcandAPIEndpoint, backend)
					if err != nil {
						log.Println(fmt.Sprintf("Error creating backend %s", err.Error()))
					}
				}

				if eBackend != nil || utf8.RuneCountInString(eBackend.Backend.ID) > 0 {

					server := new(domain.Server)
					server.Server.ID = fmt.Sprintf("%s:%d", container.ID, port.PublicPort)
					server.Server.URL = fmt.Sprintf("http://%s:%d", config.HostIP, port.PublicPort)

					log.Println(server)

					eServer, err := vulcandAPIClient.GetServerById(config.VulcandAPIEndpoint, backend.Backend.ID, server.Server.ID)
					if err != nil {
						log.Println(fmt.Sprintf("Error finding server %s", err.Error()))
					}

					if eServer == nil || utf8.RuneCountInString(eServer.Server.ID) <= 0 {
						eServer, err = vulcandAPIClient.UpsertServer(config.VulcandAPIEndpoint, backend.Backend.ID, server)
						if err != nil {
							log.Println(fmt.Sprintf("Error creating server %s", err.Error()))
						}
					}

				}
			}
		}
	}

	return nil
}
