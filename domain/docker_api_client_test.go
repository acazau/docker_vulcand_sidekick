package domain

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type _Mock_Good_DockerAPIClientManager struct{}
type _Mock_Bad_DockerAPIClientManager struct{}

func (repo *_Mock_Good_DockerAPIClientManager) ListContainers(socketPath string) ([]*Container, error) {
	var containers = []*Container{
		&Container{
			ID: "01",
		},
	}
	return containers, nil
}

func (repo *_Mock_Bad_DockerAPIClientManager) ListContainers(socketPath string) ([]*Container, error) {
	var containers []*Container
	return containers, errors.New("Bad DockerAPIClientManager")
}

func TestListContainers(t *testing.T) {
	Convey("Validate ListContainers", t, func() {

		Convey("Validate when DockerAPIClientManager is null, returns errors", func() {
			dockerAPIClient := new(DockerAPIClientManager)
			containers, err := dockerAPIClient.ListContainers("test")
			So(containers, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

		Convey("Validate when DockerAPIClientManager is valid, returns no errors", func() {
			dockerAPIClient := &DockerAPIClientManager{InjectedDockerAPIClientManager: &_Mock_Good_DockerAPIClientManager{}}
			containers, err := dockerAPIClient.ListContainers("test")
			So(len(containers), ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Validate when DockerAPIClientManager is invalid, returns errors", func() {
			dockerAPIClient := &DockerAPIClientManager{InjectedDockerAPIClientManager: &_Mock_Bad_DockerAPIClientManager{}}
			containers, err := dockerAPIClient.ListContainers("test")
			So(containers, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

	})
}
