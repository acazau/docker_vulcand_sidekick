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

type Workflow01_Repository struct {
	useCaseManager UseCaseManager
}

func init() {
	dockerAPIRepo := &infrastructure.DockerAPIClient_HTTP_Repository{}
	vulcandAPIRepo := &infrastructure.VulcandAPIClient_HTTP_Repository{}
	dockerAPIClient.InjectedDockerAPIClientManager = dockerAPIRepo
	vulcandAPIClient.InjectedVulcandAPIClientManager = vulcandAPIRepo
}

func (manager *Workflow01_Repository) Execute(arg map[string]interface{}) error {
	dockerAPIEndpoint := (arg["dockerAPIEndpoint"]).(string)
	vulcandAPIEndpoint := (arg["vulcandAPIEndpoint"]).(string)

	containers, err := dockerAPIClient.ListContainers(dockerAPIEndpoint)
	if err != nil {
		log.Println(fmt.Sprintf("Error getting containers %s", err.Error()))
	}

	for _, container := range containers {
		log.Println(fmt.Sprintf("Image %s", container.Image))
		for _, port := range container.Ports {
			if port.Type == "tcp" && port.Publicport > 0 {
				log.Println(fmt.Sprintf("Public Port %d Type %s", port.Publicport, port.Type))

				backend := new(domain.Backend)
				backend.Backend.ID = strings.Replace(container.Image, "/", "_", -1)
				backend.Backend.Type = "http"

				eBackend, err := vulcandAPIClient.GetBackendById(vulcandAPIEndpoint, backend.Backend.ID)
				if err != nil {
					log.Println(fmt.Sprintf("Error Finding backend %s", err.Error()))
				}
				if eBackend == nil || utf8.RuneCountInString(eBackend.Backend.ID) <= 0 {
					_, err := vulcandAPIClient.UpsertBackend(vulcandAPIEndpoint, backend)
					if err != nil {
						log.Println(fmt.Sprintf("Error creating backend %s", err.Error()))
						return err
					}
				}
			}
		}
	}

	return nil
}
