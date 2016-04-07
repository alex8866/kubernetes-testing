package main

import (
	"io/ioutil"
	"os/exec"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKubeCommands(t *testing.T) {
	Convey("Subject: test kubectl api-versions", t, func() {
		version, err := exec.Command("kubectl", "api-versions").Output()
		Convey("Api version v1 should be contained and no error", func() {
			So(string(version), ShouldContainSubstring, "v1")
			So(err, ShouldBeNil)
		})
	})
	Convey("Subject: test kubectl patch", t, func() {
		output, err := exec.Command("kubectl", "get", "node", "-o", "json").Output()
		So(err, ShouldBeNil)
		ioutil.WriteFile("files/node.json", output, 0644)
		output, err = exec.Command("kubectl", "patch", "-f", "files/node.json", "-p", `{"spec":{"unschedulable":true}}`).Output()
		Convey("it says the node is patched", func() {
			So(string(output), ShouldContainSubstring, "patched")
			So(err, ShouldBeNil)
		})
		output, err = exec.Command("kubectl", "get", "node").Output()
		Convey("the node is SchedulingDisabled", func() {
			So(string(output), ShouldContainSubstring, "SchedulingDisabled")
			So(err, ShouldBeNil)
		})
		output, err = exec.Command("kubectl", "patch", "-f", "files/node.json", "-p", `{"spec":{"unschedulable":false}}`).Output()
		Convey("it says the node is patched ", func() {
			So(string(output), ShouldContainSubstring, "patched")
			So(err, ShouldBeNil)
		})
		output, err = exec.Command("kubectl", "get", "node").Output()
		Convey("the node is ready, without SchedulingDisabled", func() {
			So(string(output), ShouldNotContainSubstring, "SchedulingDisabled")
			So(err, ShouldBeNil)
		})
	})
}
