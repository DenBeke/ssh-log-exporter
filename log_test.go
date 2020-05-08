package sshlogexporter

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestLogParse(t *testing.T) {
	Convey("Test extract SSH from log Line", t, func() {

		Convey("Valid SSH line", func() {
			line := "May  8 13:25:57 denbeke sshd[6002]: Invalid user pipo from 182.61.46.245 port 53214"

			output, err := ExtractSSHFromLine(line)
			So(err, ShouldBeNil)
			So(output, ShouldNotBeNil)
			So(output, ShouldHaveSameTypeAs, &SSHLogLine{})

		})

		Convey("Invalid SSH line", func() {
			line := "May  8 13:25:57 denbeke sshd[6002]"

			output, err := ExtractSSHFromLine(line)
			So(err, ShouldBeNil)
			So(output, ShouldBeNil)

		})

		Convey("Not a SSH line", func() {
			line := "May  8 13:25:56 denbeke systemd: pam_unix(systemd-user:session): session opened for user denbeke by (uid=0)"

			output, err := ExtractSSHFromLine(line)
			So(err, ShouldBeNil)
			So(output, ShouldBeNil)

		})

	})
}
