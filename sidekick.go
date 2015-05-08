package main // import "github.com/acazau/docker_vulcand_sidekick"

import (
	"flag"
	"github.com/acazau/docker_vulcand_sidekick/usecases"
	"log"
	"fmt"
)

var (
	dockerAPIEndpoint  = flag.String("d", "/var/run/docker.sock", "Docker API endpoint")
	vulcandAPIEndpoint = flag.String("v", "127.0.0.1:8182", "Vulcand API endpoint")
	workflow01Client   = new(usecases.UseCaseManager)
)

func init() {
	workflow01Repo := &usecases.Workflow01_Repository{}
	workflow01Client.InjectedUseCaseManager = workflow01Repo
}

func main() {
	map1 := map[string]interface{}{"dockerAPIEndpoint": *dockerAPIEndpoint, "vulcandAPIEndpoint": *vulcandAPIEndpoint}
	err := workflow01Client.Execute(map1)
		if err != nil {
		log.Println(fmt.Sprintf("Error getting containers %s", err.Error()))
	}
}
