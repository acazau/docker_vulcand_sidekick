package domain

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type _Mock_Good_VulcandAPIClientManager struct{}
type _Mock_Bad_VulcandAPIClientManager struct{}

func (repo *_Mock_Good_VulcandAPIClientManager) ListBackends(apiPort string) ([]*Backend, error) {
	var backends = []*Backend{
		&Backend{
			ID: "01",
		},
	}
	return backends, nil
}

func (repo *_Mock_Bad_VulcandAPIClientManager) ListBackends(apiPort string) ([]*Backend, error) {
	var backends []*Backend
	return backends, errors.New("Bad VulcandAPIClientManager")
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
