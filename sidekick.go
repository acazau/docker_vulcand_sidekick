package main // import "github.com/acazau/docker_vulcand_sidekick"

import (
	"flag"
	"github.com/acazau/docker_vulcand_sidekick/domain"
	"github.com/acazau/docker_vulcand_sidekick/infrastructure"
	"github.com/acazau/docker_vulcand_sidekick/usecases"
)

var (
	dockerAPIClient  = new(domain.DockerAPIClientManager)
	vulcandAPIClient = new(domain.VulcandAPIClientManager)
	workflow01Client = new(usecases.UseCaseManager)
	config           = &domain.Config{}
)

func init() {
	dockerAPIRepo := &infrastructure.DockerAPIClient_HTTP_Repository{}
	vulcandAPIRepo := &infrastructure.VulcandAPIClient_HTTP_Repository{}
	dockerAPIClient.InjectedDockerAPIClientManager = dockerAPIRepo
	vulcandAPIClient.InjectedVulcandAPIClientManager = vulcandAPIRepo
	workflow01Repo := &usecases.Vulcand_Register_Backend_Server_Repository{}
	workflow01Client.InjectedUseCaseManager = workflow01Repo
	config.InstallFlags()
}

func main() {
	flag.Parse()
}
