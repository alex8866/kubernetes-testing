package main

import (
	"os/exec"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKubeCluster(t *testing.T) {
	Convey("Subject: Checking kube cluster are in health", t, func() {
		Convey("Given kube cluster is configured by ansible already", func() {
			Convey("When run kubectl cluster-info", func() {
				info, err := exec.Command("kubectl", "cluster-info").Output()
				Convey("Port *443 should be contained and no error", func() {
					So(string(info), ShouldContainSubstring, "443")
					So(err, ShouldBeNil)
				})
			})
			Convey("When run kubectl get serviceaccount", func() {
				output, err := exec.Command("kubectl", "get", "serviceaccount").Output()
				Convey("ServiceAccount default should be contained and no error", func() {
					So(string(output), ShouldContainSubstring, "default")
					So(err, ShouldBeNil)
				})
			})
		})
	})
}
