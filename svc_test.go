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
		Convey("Subject: Creating service should be successfully", func() {
			Convey("Given the service example_svc is not existed yet", func() {
				Convey("When run kubectl create -f example_svc.yaml", func() {
					output, err := exec.Command("kubectl", "create", "-f", file).Output()
					Convey("it should return service example_svc created", func() {
						So(string(output), ShouldContainSubstring, "created")
						So(err, ShouldBeNil)
					})
				})
			})
		})
		Convey("Subject: Deleting service should be successfully.", func() {
			Convey("Given service example_svc is existing", func() {
				Convey("When run kubectl delete svc example_svc", func() {
					output, err := exec.Command("kubectl", "delete", "-f", file).Output()
					Convey("it should return service example_svc deleted", func() {
						So(string(output), ShouldContainSubstring, "deleted")
						So(err, ShouldBeNil)
					})
				})
			})
		})
	})
}
