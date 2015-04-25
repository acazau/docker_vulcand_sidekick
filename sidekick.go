package main // import "github.com/acazau/docker_vulcand_sidekick"

import (
	"flag"
	"github.com/acazau/docker_vulcand_sidekick/domain"
	"github.com/acazau/docker_vulcand_sidekick/infrastructure"
)

var (
	dockerAPIEndpoint  = flag.String("d", "/var/run/docker.sock", "Docker API endpoint")
	vulcandAPIEndpoint = flag.String("v", "127.0.0.1:8182", "Vulcand API endpoint")
	dockerAPIClient    = new(domain.DockerAPIClientManager)
	vulcandAPIClient   = new(domain.VulcandAPIClientManager)
)

func init() {
	dockerAPIRepo := &infrastructure.DockerAPIClient_HTTP_Repository{}
	vulcandAPIRepo := &infrastructure.VulcandAPIClient_HTTP_Repository{}
	dockerAPIClient.InjectedDockerAPIClientManager = dockerAPIRepo
	vulcandAPIClient.InjectedVulcandAPIClientManager = vulcandAPIRepo
}

func main() {
}
