package utils

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestIsValidIP4(t *testing.T) {
	Convey("Validate isValidIP4", t, func() {

		Convey("Validate ip is empty returns false", func() {
			result := IsValidIP4("")
			So(result, ShouldBeFalse)
		})

		Convey("Validate invalid returns false", func() {
			result := IsValidIP4("305.17.8.101")
			So(result, ShouldBeFalse)
		})

		Convey("Validate valid ip returns true", func() {
			result := IsValidIP4("172.17.8.101")
			So(result, ShouldBeTrue)
		})
	})
}
