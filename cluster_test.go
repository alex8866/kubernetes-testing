package main

import (
	"os/exec"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKubeCluster(t *testing.T) {
	Convey("Subject: Checking kube cluster are in health", t, func() {
		Convey("Given kube cluster is configured by ansible", func() {
			Convey("When run kubectl cluster-info", func() {
				info, err := exec.Command("kubectl", "cluster-info").Output()
				Convey("'Kubernetes master is running at' is showing", func() {
					So(string(info), ShouldContainSubstring, "running")
					So(string(info), ShouldContainSubstring, "443")
					So(err, ShouldBeNil)
				})
			})
			Convey("When run kubectl get serviceaccount", func() {
				output, err := exec.Command("kubectl", "get", "serviceaccount").Output()
				Convey("A default serviceAccount default should be listed there", func() {
					So(string(output), ShouldContainSubstring, "default")
					So(err, ShouldBeNil)
				})
			})
			Convey("When run kubectl get secret", func() {
				output, err := exec.Command("kubectl", "get", "secret").Output()
				Convey("A default token should be listed there", func() {
					So(string(output), ShouldContainSubstring, "default-token")
					So(err, ShouldBeNil)
				})
			})
			Convey("When run kubectl get namespace", func() {
				output, err := exec.Command("kubectl", "get", "namespace").Output()
				Convey("default and kube-system should be listed there", func() {
					So(string(output), ShouldContainSubstring, "default")
					So(string(output), ShouldContainSubstring, "kube-system")
					So(err, ShouldBeNil)
				})
			})
			Convey("When run kubectl get node", func() {
				output, err := exec.Command("kubectl", "get", "node").Output()
				Convey("At least one node should be ready", func() {
					So(string(output), ShouldContainSubstring, "Ready")
					So(err, ShouldBeNil)
				})
			})
			Convey("When run kubectl get service", func() {
				output, err := exec.Command("kubectl", "get", "service").Output()
				Convey("A default service named kubernetes is there", func() {
					So(string(output), ShouldContainSubstring, "kubernetes")
					So(err, ShouldBeNil)
				})
			})
		})
	})
}
