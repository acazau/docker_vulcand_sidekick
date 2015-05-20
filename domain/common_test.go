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
			So(config.DockerAPIEndpoint, ShouldNotBeEmpty)
			So(config.VulcandAPIEndpoint, ShouldNotBeEmpty)
			So(config.HostIP, ShouldNotBeEmpty)
		})
	})
}
