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

type Frontend struct {
	ID        string `json:"Id"`
	Route     string `json:"Route"`
	Type      string `json:"Type"`
	Backendid string `json:"BackendId"`
	Settings  struct {
		Limits struct {
			Maxmembodybytes int `json:"MaxMemBodyBytes"`
			Maxbodybytes    int `json:"MaxBodyBytes"`
		} `json:"Limits"`
		Failoverpredicate  string `json:"FailoverPredicate"`
		Hostname           string `json:"Hostname"`
		Trustforwardheader bool   `json:"TrustForwardHeader"`
	} `json:"Settings"`
}

type IVulcandAPIClientManager interface {
	GetBackendById(apiUrl, backendId string) (*Backend, error)
	GetServerById(apiUrl, backendId, serverId string) (*Server, error)
	GetFrontendById(apiUrl, frontendId string) (*Frontend, error)
	ListBackends(apiUrl string) ([]*Backend, error)
	ListServers(apiUrl, backendId string) ([]*Server, error)
	ListFrontends(apiUrl string) ([]*Frontend, error)
	UpsertBackend(apiUrl string, backend *Backend) (*Backend, error)
	UpsertServer(apiUrl, backendId string, server *Server) (*Server, error)
	DeleteBackendById(apiUrl, backendId string) error
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

func (manager *VulcandAPIClientManager) GetServerById(apiUrl, backendId, serverId string) (*Server, error) {
	if manager.InjectedVulcandAPIClientManager == nil {
		return nil, errors.New("Injected VulcandAPIClientManager cannot be null")
	}
	server, err := manager.InjectedVulcandAPIClientManager.GetServerById(apiUrl, backendId, serverId)

	return server, err
}

func (manager *VulcandAPIClientManager) GetFrontendById(apiUrl, frontendId string) (*Frontend, error) {
	if manager.InjectedVulcandAPIClientManager == nil {
		return nil, errors.New("Injected VulcandAPIClientManager cannot be null")
	}
	backend, err := manager.InjectedVulcandAPIClientManager.GetFrontendById(apiUrl, frontendId)

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

func (manager *VulcandAPIClientManager) DeleteBackendById(apiUrl, backendId string) error {
	if manager.InjectedVulcandAPIClientManager == nil {
		return errors.New("Injected VulcandAPIClientManager cannot be null")
	}
	err := manager.InjectedVulcandAPIClientManager.DeleteBackendById(apiUrl, backendId)

	return err
}

func (manager *VulcandAPIClientManager) ListFrontends(apiUrl string) ([]*Frontend, error) {
	if manager.InjectedVulcandAPIClientManager == nil {
		return nil, errors.New("Injected VulcandAPIClientManager cannot be null")
	}
	frontends, err := manager.InjectedVulcandAPIClientManager.ListFrontends(apiUrl)

	return frontends, err
}
