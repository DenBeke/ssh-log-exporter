package sshlogexporter

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSSHParse(t *testing.T) {
	Convey("Test SSH Parse Line", t, func() {

		testData := []struct {
			Input          string
			IsAttempt      bool
			ExpectedOutput SSHLogLine
		}{
			{
				"Accepted publickey for denbeke from 192.168.1.1 port 54822 ssh2: RSA SHA256:grzgrgkzrgkrzmgnkzrkhgzrhgrzhzr",
				false,
				SSHLogLine{
					Other,
					"192.168.1.1",
					"",
					"denbeke",
					"Accepted publickey for denbeke from 192.168.1.1 port 54822 ssh2: RSA SHA256:grzgrgkzrgkrzmgnkzrkhgzrhgrzhzr",
				},
			},
			{
				"pam_unix(sshd:session): session opened for user denbeke by (uid=42)",
				false,
				SSHLogLine{
					SessionOpened,
					"",
					"",
					"denbeke",
					"pam_unix(sshd:session): session opened for user denbeke by (uid=42)",
				},
			},
			{
				"Invalid user pipo from 182.61.46.245 port 53214",
				true,
				SSHLogLine{
					InvalidUser,
					"182.61.46.245",
					"CN",
					"pipo",
					"Invalid user pipo from 182.61.46.245 port 53214",
				},
			},
			{
				"Received disconnect from 182.61.46.245 port 53214:11: Bye Bye [preauth]",
				false,
				SSHLogLine{
					Other,
					"182.61.46.245",
					"CN",
					"",
					"Received disconnect from 182.61.46.245 port 53214:11: Bye Bye [preauth]",
				},
			},
			{
				"Disconnected from invalid user pipo 182.61.46.245 port 53214 [preauth]",
				false,
				SSHLogLine{
					Other,
					"182.61.46.245",
					"CN",
					"pipo",
					"Disconnected from invalid user pipo 182.61.46.245 port 53214 [preauth]",
				},
			},
			{
				"Connection closed by 129.158.125.138 port 8713 [preauth]",
				false,
				SSHLogLine{
					Other,
					"129.158.125.138",
					"US",
					"",
					"Connection closed by 129.158.125.138 port 8713 [preauth]",
				},
			},
			{
				"Invalid user 45.77.54.162 from 117.131.60.57 port 60269",
				true,
				SSHLogLine{
					InvalidUser,
					"117.131.60.57",
					"CN",
					"45.77.54.162",
					"Invalid user 45.77.54.162 from 117.131.60.57 port 60269",
				},
			},
			{
				"Did not receive identification string from 183.111.104.197 port 17413",
				true,
				SSHLogLine{
					DidNotReceiveIdentificationString,
					"183.111.104.197",
					"KR",
					"",
					"Did not receive identification string from 183.111.104.197 port 17413",
				},
			},
			{
				"Disconnected from authenticating user root 117.131.60.57 port 34938 [preauth]",
				true,
				SSHLogLine{
					DisconnectedFromAuthenticating,
					"117.131.60.57",
					"CN",
					"root",
					"Disconnected from authenticating user root 194.180.224.130 port 34938 [preauth]",
				},
			},
		}

		for _, t := range testData {

			Convey(t.Input, func() {

				output, err := ParseSSHLine(t.Input)
				So(err, ShouldBeNil)
				So(output.Type, ShouldEqual, t.ExpectedOutput.Type)
				So(output.IP, ShouldEqual, t.ExpectedOutput.IP)
				So(output.Country, ShouldEqual, t.ExpectedOutput.Country)
				So(output.Username, ShouldEqual, t.ExpectedOutput.Username)

				So(output.IsAttackAttempt(), ShouldEqual, t.IsAttempt)

			})

		}

	})
}
