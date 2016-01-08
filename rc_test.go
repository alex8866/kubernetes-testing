package main

import (
	"os/exec"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKubeReplicaController(t *testing.T) {
	Convey("Subject: Test replicationcontroller", t, func() {
		file := filepath.Join("files", "frontend-controller.yaml")
		Convey("Subject: Creating replicationcontroller should be successfully", func() {
			Convey("Given the replicationcontroller example_rc is not existed yet", func() {
				Convey("When run kubectl create -f example_rc.yaml", func() {
					output, err := exec.Command("kubectl", "create", "-f", file).Output()
					Convey("it should return replicationcontroller example_rc created", func() {
						So(string(output), ShouldContainSubstring, "created")
						So(err, ShouldBeNil)
					})
				})
			})
		})
		Convey("Subject: Deleting rc should be successfully.", func() {
			Convey("Given replicationcontroller example_rc is existing", func() {
				Convey("When run kubectl delete rc example_rc", func() {
					output, err := exec.Command("kubectl", "delete", "-f", file).Output()
					Convey("it should return replicationcontroller example_rc deleted", func() {
						So(string(output), ShouldContainSubstring, "deleted")
						So(err, ShouldBeNil)
					})
				})
			})
		})
	})
}
