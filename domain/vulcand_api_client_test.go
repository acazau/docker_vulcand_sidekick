package domain

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type _Mock_Good_VulcandAPIClientManager struct{}
type _Mock_Bad_VulcandAPIClientManager struct{}

func (repo *_Mock_Good_VulcandAPIClientManager) GetBackendById(apiPort, backendId string) (*Backend, error) {
	var backend = new(Backend)
	backend.Backend.ID = "01"
	backend.Backend.Type = "http"
	return backend, nil
}

func (repo *_Mock_Good_VulcandAPIClientManager) GetServerById(apiUrl, backendId, serverId string) (*Server, error) {
	var server = new(Server)
	server.Server.ID = "s01"
	server.Server.URL = "http://localhost:5001"
	server.TTL = "5s"
	return server, nil
}

func (repo *_Mock_Good_VulcandAPIClientManager) ListBackends(apiPort string) ([]*Backend, error) {
	var backend = new(Backend)
	backend.Backend.ID = "01"
	backend.Backend.Type = "http"
	var backends = []*Backend{
		backend,
	}
	return backends, nil
}

func (repo *_Mock_Good_VulcandAPIClientManager) UpsertBackend(apiUrl string, backend *Backend) (*Backend, error) {
	var uBackend = new(Backend)
	uBackend.Backend.ID = "01"
	uBackend.Backend.Type = "http"
	return uBackend, nil
}

func (repo *_Mock_Good_VulcandAPIClientManager) ListServers(apiPort, backendId string) ([]*Server, error) {
	var servers = []*Server{
		&Server{},
	}
	return servers, nil
}

func (repo *_Mock_Good_VulcandAPIClientManager) UpsertServer(apiUrl, backendId string, server *Server) (*Server, error) {
	var uServer = new(Server)
	uServer.Server.ID = "s01"
	uServer.Server.URL = "http://localhost:5001"
	uServer.TTL = "5s"
	return uServer, nil
}

func (repo *_Mock_Good_VulcandAPIClientManager) DeleteBackendById(apiUrl, backendId string) error {
	type response map[string]interface{}
	return nil
}

func (repo *_Mock_Good_VulcandAPIClientManager) ListFrontends(apiUrl string) ([]*Frontend, error) {
	var frontend = new(Frontend)
	frontend.ID = "01"
	var frontends = []*Frontend{
		frontend,
	}
	return frontends, nil
}

func (repo *_Mock_Bad_VulcandAPIClientManager) GetBackendById(apiPort, backendId string) (*Backend, error) {
	var backend Backend
	return &backend, errors.New("Bad VulcandAPIClientManager")
}

func (repo *_Mock_Bad_VulcandAPIClientManager) GetServerById(apiUrl, backendId, serverId string) (*Server, error) {
	var server Server
	return &server, errors.New("Bad VulcandAPIClientManager")
}

func (repo *_Mock_Bad_VulcandAPIClientManager) ListBackends(apiPort string) ([]*Backend, error) {
	var backends []*Backend
	return backends, errors.New("Bad VulcandAPIClientManager")
}

func (repo *_Mock_Bad_VulcandAPIClientManager) UpsertBackend(apiUrl string, backend *Backend) (*Backend, error) {
	var uBackend Backend
	return &uBackend, errors.New("Bad VulcandAPIClientManager")
}

func (repo *_Mock_Bad_VulcandAPIClientManager) ListServers(apiPort, backendId string) ([]*Server, error) {
	var servers []*Server
	return servers, errors.New("Bad VulcandAPIClientManager")
}

func (repo *_Mock_Bad_VulcandAPIClientManager) UpsertServer(apiUrl, backendId string, server *Server) (*Server, error) {
	var Server Server
	return &Server, errors.New("Bad VulcandAPIClientManager")
}

func (repo *_Mock_Bad_VulcandAPIClientManager) DeleteBackendById(apiUrl, backendId string) error {
	return errors.New("Bad VulcandAPIClientManager")
}

func (repo *_Mock_Bad_VulcandAPIClientManager) ListFrontends(apiUrl string) ([]*Frontend, error) {
	var frontends []*Frontend
	return frontends, errors.New("Bad VulcandAPIClientManager")
}

func TestGetBackendById(t *testing.T) {
	Convey("Validate GetBackendById", t, func() {

		Convey("Validate when VulcandAPIClientManager is null, returns errors", func() {
			vulcandAPIClient := new(VulcandAPIClientManager)
			backend, err := vulcandAPIClient.GetBackendById("test", "b1")
			_backend := new(Backend)
			So(backend, ShouldHaveSameTypeAs, _backend)
			So(err, ShouldNotBeNil)
		})

		Convey("Validate when VulcandAPIClientManager is valid, returns no errors", func() {
			vulcandAPIClient := &VulcandAPIClientManager{InjectedVulcandAPIClientManager: &_Mock_Good_VulcandAPIClientManager{}}
			backend, err := vulcandAPIClient.GetBackendById("test", "b1")
			_backend := new(Backend)
			So(backend, ShouldHaveSameTypeAs, _backend)
			So(err, ShouldBeNil)
		})

		Convey("Validate when VulcandAPIClientManager is invalid, returns errors", func() {
			vulcandAPIClient := &VulcandAPIClientManager{InjectedVulcandAPIClientManager: &_Mock_Bad_VulcandAPIClientManager{}}
			backend, err := vulcandAPIClient.GetBackendById("test", "b1")
			_backend := new(Backend)
			So(backend, ShouldHaveSameTypeAs, _backend)
			So(err, ShouldNotBeNil)
		})

	})
}

func TestGetServerById(t *testing.T) {
	Convey("Validate GetServerById", t, func() {

		Convey("Validate when VulcandAPIClientManager is null, returns errors", func() {
			vulcandAPIClient := new(VulcandAPIClientManager)
			server, err := vulcandAPIClient.GetServerById("test", "b1", "srv1")
			_server := new(Server)
			So(server, ShouldHaveSameTypeAs, _server)
			So(err, ShouldNotBeNil)
		})

		Convey("Validate when VulcandAPIClientManager is valid, returns no errors", func() {
			vulcandAPIClient := &VulcandAPIClientManager{InjectedVulcandAPIClientManager: &_Mock_Good_VulcandAPIClientManager{}}
			server, err := vulcandAPIClient.GetServerById("test", "b1", "srv1")
			_server := new(Server)
			So(server, ShouldHaveSameTypeAs, _server)
			So(err, ShouldBeNil)
		})

		Convey("Validate when VulcandAPIClientManager is invalid, returns errors", func() {
			vulcandAPIClient := &VulcandAPIClientManager{InjectedVulcandAPIClientManager: &_Mock_Bad_VulcandAPIClientManager{}}
			server, err := vulcandAPIClient.GetServerById("test", "b1", "srv1")
			_server := new(Server)
			So(server, ShouldHaveSameTypeAs, _server)
			So(err, ShouldNotBeNil)
		})

	})
}

func TestListBackends(t *testing.T) {
	Convey("Validate ListBackends", t, func() {

		Convey("Validate when VulcandAPIClientManager is null, returns errors", func() {
			vulcandAPIClient := new(VulcandAPIClientManager)
			backends, err := vulcandAPIClient.ListBackends("test")
			So(backends, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

		Convey("Validate when VulcandAPIClientManager is valid, returns no errors", func() {
			vulcandAPIClient := &VulcandAPIClientManager{InjectedVulcandAPIClientManager: &_Mock_Good_VulcandAPIClientManager{}}
			backends, err := vulcandAPIClient.ListBackends("test")
			So(len(backends), ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Validate when VulcandAPIClientManager is invalid, returns errors", func() {
			vulcandAPIClient := &VulcandAPIClientManager{InjectedVulcandAPIClientManager: &_Mock_Bad_VulcandAPIClientManager{}}
			backends, err := vulcandAPIClient.ListBackends("test")
			So(backends, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

	})
}

func TestUpsertBackend(t *testing.T) {
	var _backend = new(Backend)
	_backend.Backend.ID = "01"
	_backend.Backend.Type = "http"

	Convey("Validate UpsertBackend", t, func() {

		Convey("Validate when VulcandAPIClientManager is null, returns errors", func() {
			vulcandAPIClient := new(VulcandAPIClientManager)
			backend, err := vulcandAPIClient.UpsertBackend("test", _backend)
			So(backend, ShouldHaveSameTypeAs, _backend)
			So(err, ShouldNotBeNil)
		})

		Convey("Validate when VulcandAPIClientManager is valid, returns no errors", func() {
			vulcandAPIClient := &VulcandAPIClientManager{InjectedVulcandAPIClientManager: &_Mock_Good_VulcandAPIClientManager{}}
			backend, err := vulcandAPIClient.UpsertBackend("test", _backend)
			So(backend, ShouldHaveSameTypeAs, _backend)
			So(err, ShouldBeNil)
		})

		Convey("Validate when VulcandAPIClientManager is invalid, returns errors", func() {
			vulcandAPIClient := &VulcandAPIClientManager{InjectedVulcandAPIClientManager: &_Mock_Bad_VulcandAPIClientManager{}}
			backend, err := vulcandAPIClient.UpsertBackend("test", _backend)
			So(backend, ShouldHaveSameTypeAs, _backend)
			So(err, ShouldNotBeNil)
		})

	})
}

func TestListServers(t *testing.T) {
	Convey("Validate ListServers", t, func() {

		Convey("Validate when VulcandAPIClientManager is null, returns errors", func() {
			vulcandAPIClient := new(VulcandAPIClientManager)
			servers, err := vulcandAPIClient.ListServers("test", "b1")
			So(servers, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

		Convey("Validate when VulcandAPIClientManager is valid, returns no errors", func() {
			vulcandAPIClient := &VulcandAPIClientManager{InjectedVulcandAPIClientManager: &_Mock_Good_VulcandAPIClientManager{}}
			servers, err := vulcandAPIClient.ListServers("test", "b1")
			So(len(servers), ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Validate when VulcandAPIClientManager is invalid, returns errors", func() {
			vulcandAPIClient := &VulcandAPIClientManager{InjectedVulcandAPIClientManager: &_Mock_Bad_VulcandAPIClientManager{}}
			servers, err := vulcandAPIClient.ListServers("test", "b1")
			So(servers, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

	})
}

func TestUpsertServer(t *testing.T) {
	var _server = new(Server)
	_server.Server.ID = "s01"
	_server.Server.URL = "http://localhost:5001"
	_server.TTL = "5s"

	Convey("Validate UpsertServer", t, func() {

		Convey("Validate when VulcandAPIClientManager is null, returns errors", func() {
			vulcandAPIClient := new(VulcandAPIClientManager)
			server, err := vulcandAPIClient.UpsertServer("test", "b1", _server)
			So(server, ShouldHaveSameTypeAs, _server)
			So(err, ShouldNotBeNil)
		})

		Convey("Validate when VulcandAPIClientManager is valid, returns no errors", func() {
			vulcandAPIClient := &VulcandAPIClientManager{InjectedVulcandAPIClientManager: &_Mock_Good_VulcandAPIClientManager{}}
			server, err := vulcandAPIClient.UpsertServer("test", "b1", _server)
			So(server, ShouldHaveSameTypeAs, _server)
			So(err, ShouldBeNil)
		})

		Convey("Validate when VulcandAPIClientManager is invalid, returns errors", func() {
			vulcandAPIClient := &VulcandAPIClientManager{InjectedVulcandAPIClientManager: &_Mock_Bad_VulcandAPIClientManager{}}
			server, err := vulcandAPIClient.UpsertServer("test", "b1", _server)
			So(server, ShouldHaveSameTypeAs, _server)
			So(err, ShouldNotBeNil)
		})

	})
}

func TestDeleteBackendById(t *testing.T) {
	Convey("Validate DeleteBackendById", t, func() {

		Convey("Validate when VulcandAPIClientManager is null, returns errors", func() {
			vulcandAPIClient := new(VulcandAPIClientManager)
			err := vulcandAPIClient.DeleteBackendById("test", "b1")
			So(err, ShouldNotBeNil)
		})

		Convey("Validate when VulcandAPIClientManager is valid, returns no errors", func() {
			vulcandAPIClient := &VulcandAPIClientManager{InjectedVulcandAPIClientManager: &_Mock_Good_VulcandAPIClientManager{}}
			err := vulcandAPIClient.DeleteBackendById("test", "b1")
			So(err, ShouldBeNil)
		})

		Convey("Validate when VulcandAPIClientManager is invalid, returns errors", func() {
			vulcandAPIClient := &VulcandAPIClientManager{InjectedVulcandAPIClientManager: &_Mock_Bad_VulcandAPIClientManager{}}
			err := vulcandAPIClient.DeleteBackendById("test", "b1")
			So(err, ShouldNotBeNil)
		})

	})
}

func TestListFrontends(t *testing.T) {
	Convey("Validate ListFrontends", t, func() {

		Convey("Validate when VulcandAPIClientManager is null, returns errors", func() {
			vulcandAPIClient := new(VulcandAPIClientManager)
			frontends, err := vulcandAPIClient.ListFrontends("test")
			So(frontends, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

		Convey("Validate when VulcandAPIClientManager is valid, returns no errors", func() {
			vulcandAPIClient := &VulcandAPIClientManager{InjectedVulcandAPIClientManager: &_Mock_Good_VulcandAPIClientManager{}}
			frontends, err := vulcandAPIClient.ListFrontends("test")
			So(len(frontends), ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Validate when VulcandAPIClientManager is invalid, returns errors", func() {
			vulcandAPIClient := &VulcandAPIClientManager{InjectedVulcandAPIClientManager: &_Mock_Bad_VulcandAPIClientManager{}}
			frontends, err := vulcandAPIClient.ListFrontends("test")
			So(frontends, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

	})
}
