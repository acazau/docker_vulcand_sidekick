package main // import "github.com/acazau/docker_vulcand_sidekick"

import (
	"flag"
	"github.com/acazau/docker_vulcand_sidekick/domain"
	"github.com/acazau/docker_vulcand_sidekick/infrastructure"
)

var (
	dockerAPIClient  = new(domain.DockerAPIClientManager)
	vulcandAPIClient = new(domain.VulcandAPIClientManager)
	config           = &domain.Config{}
)

func init() {
	dockerAPIRepo := &infrastructure.DockerAPIClient_HTTP_Repository{}
	vulcandAPIRepo := &infrastructure.VulcandAPIClient_HTTP_Repository{}
	dockerAPIClient.InjectedDockerAPIClientManager = dockerAPIRepo
	vulcandAPIClient.InjectedVulcandAPIClientManager = vulcandAPIRepo
	config.InstallFlags()
}

func main() {
	flag.Parse()
}
