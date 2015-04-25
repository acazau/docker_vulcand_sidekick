package infrastructure

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestListBackends(t *testing.T) {
	Convey("Validate ListBackends", t, func() {

		Convey("Validate when BackendAPIClientManager is null, returns errors", func() {
			vulcandapiclient := &VulcandAPIClient_HTTP_Repository{}
			backends, err := vulcandapiclient.ListBackends("test")
			So(backends, ShouldBeEmpty)
			So(err, ShouldNotBeNil)
		})

	})
}
