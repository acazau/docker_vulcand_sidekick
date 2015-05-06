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

func (repo *_Mock_Bad_VulcandAPIClientManager) GetBackendById(apiPort, backendId string) (*Backend, error) {
	var backend Backend
	return &backend, errors.New("Bad VulcandAPIClientManager")
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
