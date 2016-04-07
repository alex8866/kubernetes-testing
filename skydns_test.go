package main

import (
	"os/exec"
	"testing"
	"strings"

	. "github.com/smartystreets/goconvey/convey"
)

func TestKubeDNS(t *testing.T) {
	Convey("Subject: Checking DNS are working", t, func() {
		Convey("Subject: Checking whether kube-dns rc is existed", func() {
				Convey("When run kubectl get --namespace=kube-system rc kube-dns-v11", func() {
					output, err := exec.Command("kubectl", "get", "--namespace=kube-system", "rc", "kube-dns-v11").Output()
					Convey("etcd, kube2sky, skydns and healthz should be in the output", func() {
						So(string(output), ShouldContainSubstring, "etcd")
						So(string(output), ShouldContainSubstring, "kube2sky")
						So(string(output), ShouldContainSubstring, "skydns")
						So(string(output), ShouldContainSubstring, "healthz")
						So(err, ShouldBeNil)
					})
				})
		})
		Convey("Subject: Checking whether all pods are running ", func() {
				Convey("When run kubectl get --all-namespaces pod", func() {
					output, err := exec.Command("kubectl", "get", "--namespace=kube-system", "-l", "k8s-app=kube-dns", "pod").Output()
					Convey("kube-dns-v11-xxxxx and 4/4 should be in the output", func() {
						So(string(output), ShouldContainSubstring, "kube-dns-v11")
						So(string(output), ShouldContainSubstring, "Running")
						So(string(output), ShouldContainSubstring, "4/4")
						So(err, ShouldBeNil)
					})
				})
		})
		Convey("Subject: Checking whether kube-dns svc is existed", func() {
				Convey("When run kubectl get --namespace=kube-system svc kube-dns", func() {
					output, err := exec.Command("kubectl", "get", "--namespace=kube-system", "svc", "kube-dns").Output()
					Convey("kube-dns should be in the output", func() {
						So(string(output), ShouldContainSubstring, "kube-dns")
						So(err, ShouldBeNil)
					})
				})
				Convey("When run kubectl exec --namespace=kube-system -it kube-dns-v11-xxxxx -c skydns nslookup kube-dns.kube-system.svc.cluster.local 127.0.0.1", func() {
					output, err := exec.Command("kubectl", "get", "--namespace=kube-system", "-l", "k8s-app=kube-dns", "pod").Output()
					n := strings.Index(string(output), "kube-dns-v11")
					l := strings.Count("kube-dns-v11-xxxx", "")
					podname := string(output)[n:n+l]
					output, err = exec.Command("kubectl", "exec", "--namespace=kube-system", podname, "-c", "skydns", "nslookup", "kube-dns.kube-system.svc.cluster.local", "127.0.0.1").Output()
					Convey("kube-dns and its IP 10.254.0.10 should be in the output", func() {
						So(string(output), ShouldContainSubstring, "kube-dns")
						So(string(output), ShouldContainSubstring, "10.254.0.10")
						So(err, ShouldBeNil)
					})
				})
		})
	})
}
