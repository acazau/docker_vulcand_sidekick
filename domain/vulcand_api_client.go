package domain

import (
	"errors"
)

type Backend struct {
	ID       string `json:"Id"`
	Type     string `json:"Type"`
	Settings struct {
		Timeouts struct {
			Read         string `json:"Read"`
			Dial         string `json:"Dial"`
			Tlshandshake string `json:"TLSHandshake"`
		} `json:"Timeouts"`
		Keepalive struct {
			Period              string `json:"Period"`
			Maxidleconnsperhost int    `json:"MaxIdleConnsPerHost"`
		} `json:"KeepAlive"`
	} `json:"Settings"`
}

type Server struct {
	ID  string `json:"Id"`
	URL string `json:"URL"`
}

type IVulcandAPIClientManager interface {
	ListBackends(apiUrl string) ([]*Backend, error)
	ListServers(apiUrl, backendId string) ([]*Server, error)
}

type VulcandAPIClientManager struct {
	InjectedVulcandAPIClientManager IVulcandAPIClientManager
}

func (manager *VulcandAPIClientManager) ListBackends(apiUrl string) ([]*Backend, error) {
	if manager.InjectedVulcandAPIClientManager == nil {
		return nil, errors.New("Injected VulcandAPIClientManager cannot be null")
	}
	backends, err := manager.InjectedVulcandAPIClientManager.ListBackends(apiUrl)

	return backends, err
}

func (manager *VulcandAPIClientManager) ListServers(apiUrl, backendId string) ([]*Server, error) {
	if manager.InjectedVulcandAPIClientManager == nil {
		return nil, errors.New("Injected VulcandAPIClientManager cannot be null")
	}
	servers, err := manager.InjectedVulcandAPIClientManager.ListServers(apiUrl, backendId)

	return servers, err
}
