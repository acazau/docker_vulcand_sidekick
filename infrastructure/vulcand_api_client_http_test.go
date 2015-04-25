package infrastructure

import (
	"github.com/acazau/docker_vulcand_sidekick/domain"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestGetBackendById(t *testing.T) {
	Convey("Validate GetBackendById", t, func() {

		Convey("Validate when VulcandAPIClientManager is null, returns errors", func() {
			vulcandAPIClient := &VulcandAPIClient_HTTP_Repository{}
			backend, err := vulcandAPIClient.GetBackendById("test", "b1")
			_backend := new(domain.Backend)
			So(backend, ShouldHaveSameTypeAs, _backend)
			So(err, ShouldNotBeNil)
		})

	})
}

func TestListBackends(t *testing.T) {
	Convey("Validate ListBackends", t, func() {

		Convey("Validate when VulcandAPIClientManager is null, returns errors", func() {
			vulcandAPIClient := &VulcandAPIClient_HTTP_Repository{}
			backends, err := vulcandAPIClient.ListBackends("test")
			So(backends, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

	})
}

func TestListServers(t *testing.T) {
	Convey("Validate ListServers", t, func() {

		Convey("Validate when VulcandAPIClientManager is null, returns errors", func() {
			vulcandAPIClient := &VulcandAPIClient_HTTP_Repository{}
			servers, err := vulcandAPIClient.ListServers("test", "b1")
			So(servers, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

	})
}
