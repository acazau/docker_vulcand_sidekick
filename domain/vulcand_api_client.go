package domain

import (
	"errors"
)

type Backend struct {
	Backend struct {
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
	} `json:"Backend"`
}

type Server struct {
	Server struct {
		ID  string `json:"Id"`
		URL string `json:"URL"`
	} `json:"Server"`
	TTL string `json:"TTL"`
}

type IVulcandAPIClientManager interface {
	GetBackendById(apiUrl, backendId string) (*Backend, error)
	ListBackends(apiUrl string) ([]*Backend, error)
	ListServers(apiUrl, backendId string) ([]*Server, error)
	UpsertBackend(apiUrl string, backend *Backend) (*Backend, error)
	UpsertServer(apiUrl, backendId string, server *Server) (*Server, error)
}

type VulcandAPIClientManager struct {
	InjectedVulcandAPIClientManager IVulcandAPIClientManager
}

func (manager *VulcandAPIClientManager) GetBackendById(apiUrl, backendId string) (*Backend, error) {
	if manager.InjectedVulcandAPIClientManager == nil {
		return nil, errors.New("Injected VulcandAPIClientManager cannot be null")
	}
	backend, err := manager.InjectedVulcandAPIClientManager.GetBackendById(apiUrl, backendId)

	return backend, err
}

func (manager *VulcandAPIClientManager) ListBackends(apiUrl string) ([]*Backend, error) {
	if manager.InjectedVulcandAPIClientManager == nil {
		return nil, errors.New("Injected VulcandAPIClientManager cannot be null")
	}
	backends, err := manager.InjectedVulcandAPIClientManager.ListBackends(apiUrl)

	return backends, err
}

func (manager *VulcandAPIClientManager) UpsertBackend(apiUrl string, backend *Backend) (*Backend, error) {
	if manager.InjectedVulcandAPIClientManager == nil {
		return nil, errors.New("Injected VulcandAPIClientManager cannot be null")
	}
	upsertedBackend, err := manager.InjectedVulcandAPIClientManager.UpsertBackend(apiUrl, backend)

	return upsertedBackend, err
}

func (manager *VulcandAPIClientManager) ListServers(apiUrl, backendId string) ([]*Server, error) {
	if manager.InjectedVulcandAPIClientManager == nil {
		return nil, errors.New("Injected VulcandAPIClientManager cannot be null")
	}
	servers, err := manager.InjectedVulcandAPIClientManager.ListServers(apiUrl, backendId)

	return servers, err
}

func (manager *VulcandAPIClientManager) UpsertServer(apiUrl, backendId string, server *Server) (*Server, error) {
	if manager.InjectedVulcandAPIClientManager == nil {
		return nil, errors.New("Injected VulcandAPIClientManager cannot be null")
	}
	upsertedServer, err := manager.InjectedVulcandAPIClientManager.UpsertServer(apiUrl, backendId, server)

	return upsertedServer, err
}
