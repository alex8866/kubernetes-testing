package main

import (
	"bytes"
	"os/exec"
	"strconv"
	"strings"
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

type version struct {
	Major        string
	Minor        string
	GitVersion   string
	GitCommit    string
	GitTreeState string
}

type Version struct {
	V map[string]map[string]string
}

func TestKubeVersion(t *testing.T) {
	Convey("Subject: Version Checking", t, func() {
		kversion, err := exec.Command("kubectl", "version").Output()
		Convey("kubectl version returns no error", func() {
			So(err, ShouldBeNil)
		})
		versions := bytes.Split(kversion, []byte("\n"))
		cver := string(versions[0])
		cver = cver[strings.IndexRune(cver, '{')+1:]
		cver = strings.TrimRight(cver, "}")
		sver := string(versions[1])
		sver = sver[strings.IndexRune(sver, '{')+1:]
		sver = strings.TrimRight(sver, "}")
		Convey("Subject: Client and Server version should be equal", func() {
			So(cver, ShouldEqual, sver)
		})
		kver := strings.Split(cver, ",")
		vmap := make(map[string]string)
		for _, item := range kver {
			entry := strings.Split(strings.TrimSpace(item), ":")
			vmap[entry[0]] = entry[1]
		}

		rversion, err := exec.Command("rpm", "-q", "kubernetes").Output()
		Convey("rpm -q kubernetes returns no error", func() {
			So(err, ShouldBeNil)
		})
		srver := string(rversion)
		rgit := strings.Split(srver, "-")[1]
		rver := strings.FieldsFunc(srver, func(r rune) bool {
			return r == '-' || r == '.'
		})
		rmajor := rver[1]
		rminor := rver[2]
		rcommit := rver[7][3:]
		vmajor, err := strconv.Unquote(vmap["Major"])
		Convey("unquote major return no error", func() {
			So(err, ShouldBeNil)
		})
		Convey("version major should match", func() {
			So(vmajor, ShouldEqual, rmajor)
		})
		vminor, err := strconv.Unquote(vmap["Minor"])
		Convey("unquote minor return no error", func() {
			So(err, ShouldBeNil)
		})
		Convey("version minor should match", func() {
			So(vminor, ShouldEqual, rminor)
		})
		Convey("rpm git version should be within kubectl git version", func() {
			So(vmap["GitVersion"], ShouldContainSubstring, rgit)
		})
		Convey("rpm git commit should be within kubectl git commit", func() {
			So(vmap["GitCommit"], ShouldContainSubstring, rcommit)
		})
	})
}
