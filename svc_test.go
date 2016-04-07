package main

import (
	"os/exec"
	"path/filepath"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKubeServices(t *testing.T) {
	Convey("Subject: Test kube services", t, func() {
		file := filepath.Join("files", "frontend-service.yaml")
		Convey("Subject: Creating service should be successful", func() {
			Convey("Given the service 'frontend' is not existed yet", func() {
				Convey("When run kubectl create -f frontend-service.yaml", func() {
					output, err := exec.Command("kubectl", "create", "-f", file).Output()
					Convey("it should return service 'frontend' created", func() {
						So(string(output), ShouldContainSubstring, "created")
						So(string(output), ShouldContainSubstring, "frontend")
						So(err, ShouldBeNil)
					})
				})
			})
		})
		Convey("Subject: Deleting service should be successful.", func() {
			Convey("Given service 'frontend' is existing", func() {
				Convey("When run kubectl delete svc frontend", func() {
					output, err := exec.Command("kubectl", "delete", "-f", file).Output()
					Convey("it should return service 'frontend' deleted", func() {
						So(string(output), ShouldContainSubstring, "deleted")
						So(string(output), ShouldContainSubstring, "frontend")
						So(err, ShouldBeNil)
					})
				})
			})
		})
	})
}
