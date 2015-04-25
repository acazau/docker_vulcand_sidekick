package infrastructure

import (
	"encoding/json"
	"fmt"
	"github.com/acazau/docker_vulcand_sidekick/domain"
	"io/ioutil"
	"log"
	"net"
	"net/http"
	"net/http/httputil"
)

type VulcandAPIClient_HTTP_Repository struct {
	client domain.VulcandAPIClientManager
}

func ExecuteRequest(apiUrl, apiQuery string) ([]byte, error) {
	conn, err := net.Dial("tcp", apiUrl)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	c := httputil.NewClientConn(conn, nil)
	defer c.Close()

	req, err := http.NewRequest("GET", apiQuery, nil)

	res, err := c.Do(req)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer res.Body.Close()

	payload, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	return payload, nil
}

func (repo *VulcandAPIClient_HTTP_Repository) GetBackendById(apiUrl, backendId string) (*domain.Backend, error) {
	apiQuery := fmt.Sprintf("/v2/backends/%s", backendId)
	payload, err := ExecuteRequest(apiUrl, apiQuery)

	backend := domain.Backend{}
	err = json.Unmarshal(payload, &backend)
	if err != nil {
		return nil, err
	}

	return &backend, nil
}

func (repo *VulcandAPIClient_HTTP_Repository) ListBackends(apiUrl string) ([]*domain.Backend, error) {
	payload, err := ExecuteRequest(apiUrl, "/v2/backends")

	var payloadUnmarshalled map[string][]*domain.Backend
	err = json.Unmarshal(payload, &payloadUnmarshalled)
	if err != nil {
		return nil, err
	}

	backends := []*domain.Backend{}
	for i := range payloadUnmarshalled["Backends"] {
		item := payloadUnmarshalled["Backends"][i]
		backends = append(backends, item)
	}

	return backends, nil
}

func (repo *VulcandAPIClient_HTTP_Repository) ListServers(apiUrl, backendId string) ([]*domain.Server, error) {
	apiQuery := fmt.Sprintf("/v2/backends/%s/servers", backendId)
	payload, err := ExecuteRequest(apiUrl, apiQuery)

	var payloadUnmarshalled map[string][]*domain.Server
	err = json.Unmarshal(payload, &payloadUnmarshalled)
	if err != nil {
		return nil, err
	}

	servers := []*domain.Server{}
	for i := range payloadUnmarshalled["Servers"] {
		item := payloadUnmarshalled["Servers"][i]
		servers = append(servers, item)
	}

	return servers, nil
}
