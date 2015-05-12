package domain

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestInstallFlags(t *testing.T) {
	Convey("Validate InstallFlags", t, func() {

		Convey("Validate InstallFlags returns default values", func() {
			config := new(Config)
			config.InstallFlags()
			So(config.DockerAPIEndpoint, ShouldEqual, "/var/run/docker.sock")
			So(config.VulcandAPIEndpoint, ShouldEqual, "127.0.0.1:8182")
			So(config.HostIP, ShouldEqual, "172.17.8.101")
		})
	})
}
