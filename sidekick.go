package main // import "github.com/acazau/docker_vulcand_sidekick"

import (
	"flag"
	"fmt"
	"github.com/acazau/docker_vulcand_sidekick/domain"
	"github.com/acazau/docker_vulcand_sidekick/infrastructure"
	"log"
)

var (
	dockerapiendpoint = flag.String("e", "/var/run/docker.sock", "Docker API endpoint")
	dockerapiclient   = new(domain.DockerAPIClientManager)
)

func init() {
	dockerApiRepo := &infrastructure.DockerAPIClientRepository{}
	dockerapiclient.InjectedDockerAPIClientManager = dockerApiRepo
}

func main() {
	_, err := dockerapiclient.ListContainers(*dockerapiendpoint)
	if err != nil {
		log.Println(fmt.Sprintf("Error getting containers %s", err.Error()))
	}
}
