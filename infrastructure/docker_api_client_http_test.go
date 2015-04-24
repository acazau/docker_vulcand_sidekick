package infrastructure

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestListContainers(t *testing.T) {
	Convey("Validate ListContainers", t, func() {

		Convey("Validate when DockerAPIClientManager is null, returns errors", func() {
			dockerapiclient := &DockerAPIClient_HTTP_Repository{}
			containers, err := dockerapiclient.ListContainers("test")
			So(containers, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

	})
}
