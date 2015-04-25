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

type VulcandAPIClient_HTTP_Repository struct {
	client domain.VulcandAPIClientManager
}

func (repo *VulcandAPIClient_HTTP_Repository) ListBackends(apiUrl string) ([]*domain.Backend, error) {
	conn, err := net.Dial("tcp", apiUrl)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	c := httputil.NewClientConn(conn, nil)
	defer c.Close()

	req, err := http.NewRequest("GET", "/v2/backends", nil)

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

	var payload map[string][]*domain.Backend
	err = json.Unmarshal(data, &payload)
	if err != nil {
		return nil, err
	}

	backends := []*domain.Backend{}
	for i := range payload["Backends"] {
		item := payload["Backends"][i]
		backends = append(backends, item)
	}

	return backends, nil
}
