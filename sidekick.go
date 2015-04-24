package main // import "github.com/acazau/docker_vulcand_sidekick"

import (
	"flag"
	"fmt"
	"github.com/acazau/docker_vulcand_sidekick/domain"
	"github.com/acazau/docker_vulcand_sidekick/infrastructure"
	"log"
)

var (
	dockerAPIEndpoint = flag.String("e", "/var/run/docker.sock", "Docker API endpoint")
	dockerAPIClient   = new(domain.DockerAPIClientManager)
)

func init() {
	dockerAPIRepo := &infrastructure.DockerAPIClient_HTTP_Repository{}
	dockerAPIClient.InjectedDockerAPIClientManager = dockerAPIRepo
}

func main() {
	_, err := dockerAPIClient.ListContainers(*dockerAPIEndpoint)
	if err != nil {
		log.Println(fmt.Sprintf("Error getting containers %s", err.Error()))
	}
}
