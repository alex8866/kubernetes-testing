package main

import (
	"os/exec"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKubeCommands(t *testing.T) {
	Convey("When run kubectl api-versions", t, func() {
		version, err := exec.Command("kubectl", "api-versions").Output()
		Convey("Api version v1 should be contained and no error", func() {
			So(string(version), ShouldContainSubstring, "v1")
			So(err, ShouldBeNil)
		})
	})
}
