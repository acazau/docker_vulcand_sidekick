package domain

import (
	"errors"
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

type _GoodDockerAPIClientManager struct{}
type _BadDockerAPIClientManager struct{}

func (repo *_GoodDockerAPIClientManager) ListContainers(socketPath string) ([]*Container, error) {
	var containers = []*Container{
		&Container{
			ID: "01",
		},
	}
	return containers, nil
}

func (repo *_BadDockerAPIClientManager) ListContainers(socketPath string) ([]*Container, error) {
	var containers []*Container
	return containers, errors.New("Bad DockerAPIClientManager")
}

func TestListContainers(t *testing.T) {
	Convey("Validate ListContainers", t, func() {

		Convey("Validate when DockerAPIClientManager is null, returns errors", func() {
			dockerapiclient := new(DockerAPIClientManager)
			containers, err := dockerapiclient.ListContainers("test")
			So(containers, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

		Convey("Validate when DockerAPIClientManager is valid, returns no errors", func() {
			dockerapiclient := &DockerAPIClientManager{InjectedDockerAPIClientManager: &_GoodDockerAPIClientManager{}}
			containers, err := dockerapiclient.ListContainers("test")
			So(len(containers), ShouldEqual, 1)
			So(err, ShouldBeNil)
		})

		Convey("Validate when DockerAPIClientManager is invalid, returns errors", func() {
			dockerapiclient := &DockerAPIClientManager{InjectedDockerAPIClientManager: &_BadDockerAPIClientManager{}}
			containers, err := dockerapiclient.ListContainers("test")
			So(containers, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

	})
}
