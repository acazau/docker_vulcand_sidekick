package infrastructure

import (
	"encoding/json"
	"github.com/acazau/docker_vulcand_sidekick/domain"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
)

type DockerAPIClient_HTTP_Repository struct {
	client domain.DockerAPIClientManager
}

func (repo *DockerAPIClient_HTTP_Repository) ListContainers(socketPath string) ([]*domain.Container, error) {
	conn, err := net.Dial("unix", socketPath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	c := httputil.NewClientConn(conn, nil)
	defer c.Close()

	req, err := http.NewRequest("GET", "/containers/json?all=1", nil)

	res, err := c.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	data, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	containers := []*domain.Container{}
	err = json.Unmarshal(data, &containers)
	if err != nil {
		return nil, err
	}

	return containers, nil
}
