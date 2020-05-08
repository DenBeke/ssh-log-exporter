package sshlogexporter

import (
	"strings"
)

const (
	sshPrefix = "sshd"
)

func IsSSH(line string) bool {
	return strings.Contains(line, sshPrefix)
}

func ExtractSSHFromLine(line string) (*SSHLogLine, error) {
	if !IsSSH(line) {
		return nil, nil
	}

	s := strings.SplitN(line, ": ", 2)

	if !(len(s) > 1) {
		return nil, nil
	}

	return ParseSSHLine(s[1])
}
