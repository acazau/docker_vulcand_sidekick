package infrastructure

import (
	"bytes"
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

func ExecuteRequest(method, apiUrl, apiQuery string, body []byte, headers map[string]string) ([]byte, error) {
	conn, err := net.Dial("tcp", apiUrl)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	c := httputil.NewClientConn(conn, nil)
	defer c.Close()

	b := bytes.NewBuffer(body)
	req, err := http.NewRequest(method, apiQuery, b)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Content-Type", "application/json")
	if headers != nil {
		for header, value := range headers {
			req.Header.Add(header, value)
		}
	}

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
	payload, err := ExecuteRequest("GET", apiUrl, apiQuery, nil, nil)

	backend := domain.Backend{}
	err = json.Unmarshal(payload, &backend)
	if err != nil {
		return nil, err
	}

	return &backend, nil
}

func (repo *VulcandAPIClient_HTTP_Repository) ListBackends(apiUrl string) ([]*domain.Backend, error) {
	payload, err := ExecuteRequest("GET", apiUrl, "/v2/backends", nil, nil)

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

func (repo *VulcandAPIClient_HTTP_Repository) UpsertBackend(apiUrl string, backend *domain.Backend) (*domain.Backend, error) {
	data, err := json.Marshal(&backend)
	if err != nil {
		return nil, err
	}
	payload, err := ExecuteRequest("POST", apiUrl, "/v2/backends", data, nil)

	upsertedBackend := domain.Backend{}
	err = json.Unmarshal(payload, &upsertedBackend)
	if err != nil {
		return nil, err
	}

	return &upsertedBackend, nil
}

func (repo *VulcandAPIClient_HTTP_Repository) ListServers(apiUrl, backendId string) ([]*domain.Server, error) {
	apiQuery := fmt.Sprintf("/v2/backends/%s/servers", backendId)
	payload, err := ExecuteRequest("GET", apiUrl, apiQuery, nil, nil)

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

func (repo *VulcandAPIClient_HTTP_Repository) UpsertServer(apiUrl, backendId string, server *domain.Server) (*domain.Server, error) {
	apiQuery := fmt.Sprintf("/v2/backends/%s/servers", backendId)
	data, err := json.Marshal(&server)
	if err != nil {
		return nil, err
	}
	payload, err := ExecuteRequest("POST", apiUrl, apiQuery, data, nil)

	upsertedServer := domain.Server{}
	err = json.Unmarshal(payload, &upsertedServer)
	if err != nil {
		return nil, err
	}

	return &upsertedServer, nil
}
