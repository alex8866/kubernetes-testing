package main

import (
	"os/exec"
	"path/filepath"
	"testing"
	"time"
	"strings"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKubeReplicaController(t *testing.T) {
	Convey("Subject: Test replicationcontroller", t, func() {
		file := filepath.Join("files", "frontend-controller.yaml")
		Convey("Subject: Creating replicationcontroller should be successful", func() {
			Convey("Given the replicationcontroller 'frontend' is not existed yet", func() {
				Convey("When run kubectl create -f frontend-controller.yaml", func() {
					output, err := exec.Command("kubectl", "create", "-f", file).Output()
					time.Sleep(time.Second * 30)
					Convey("it should return replicationcontroller 'frontend' created", func() {
						So(string(output), ShouldContainSubstring, "created")
						So(string(output), ShouldContainSubstring, "frontend")
						So(err, ShouldBeNil)
					})
				})
			})
		})
		Convey("Subject: rc scale is successful", func() {
			Convey("When run kubectl get -l tier=frontend po", func() {
				output, err := exec.Command("kubectl", "get", "-l", "tier=frontend", "pod").Output()
				Convey("it should return 3 running frontend pod", func() {
					podNum := strings.Count(string(output), "Running")
					So(podNum, ShouldEqual, 3) 
					So(err, ShouldBeNil)
				})
			})
			Convey("When run kubectl scale --replicas=4 rc/frontend", func() {
				output, err := exec.Command("kubectl", "scale", "--replicas=4", "rc/frontend").Output()
				time.Sleep(time.Second * 30)
				Convey("it should return replicationcontroller 'frontend' created", func() {
					So(string(output), ShouldContainSubstring, "scaled")
					So(string(output), ShouldContainSubstring, "frontend")
					So(err, ShouldBeNil)
				})
				output, err = exec.Command("kubectl", "get", "-l", "tier=frontend", "pod").Output()
				Convey("it should return 4 running frontend pod", func() {
					podNum := strings.Count(string(output), "Running")
					So(podNum, ShouldEqual, 4) 
					So(err, ShouldBeNil)
				})
			})
		})
		Convey("Subject: Deleting rc should be successful.", func() {
			Convey("Given replicationcontroller 'frontend' is existing", func() {
				Convey("When run kubectl delete rc frontend", func() {
					output, err := exec.Command("kubectl", "delete", "-f", file).Output()
					Convey("it should return replicationcontroller 'frontend' deleted", func() {
						So(string(output), ShouldContainSubstring, "deleted")
						So(string(output), ShouldContainSubstring, "frontend")
						So(err, ShouldBeNil)
					})
				})
			})
		})
	})
}
