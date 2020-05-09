package sshlogexporter

import (
	"net"
	"strings"

	"github.com/pariz/gountries"
	"github.com/phuslu/geoip"
)

const (
	InvalidUser                       = "InvalidUSer"
	DidNotReceiveIdentificationString = "DidNotReceiveIdentificationString"
	DisconnectedFromAuthenticating    = "DisconnectedFromAuthenticating"
	SessionOpened                     = "SessionOpened"
	Other                             = "Other"
)

type SSHLogLine struct {
	Type        string
	IP          string
	Country     string
	CountryName string
	Username    string
	Log         string
}

func (s *SSHLogLine) IsAttackAttempt() bool {
	if s.Type == SessionOpened || s.Type == Other {
		return false
	}
	return true
}

func ParseSSHLine(line string) (s *SSHLogLine, err error) {

	s = &SSHLogLine{
		Log:  line,
		Type: Other,
	}

	// Find Username and IP
	words := strings.Split(line, " ")
	for i, word := range words {

		if !(i+1 < len(words)) {
			break // no next part left in the slice
		}

		// look "for {username}"
		if word == "for" && words[i+1] == "user" {
			if i+2 < len(words) {
				s.Username = words[i+2]
				continue
			}

			s.Username = words[i+1]
			continue
		}

		if word == "for" {
			s.Username = words[i+1]
			continue
		}

		// look "for user {username}"
		if word == "user" {
			s.Username = words[i+1]
			continue
		}

		// look for IP
		if word == "port" {
			s.IP = words[i-1]
		}

	}

	s.Country = string(geoip.Country(net.ParseIP(s.IP)))
	if s.Country == "ZZ" {
		s.Country = ""
	}

	if s.Country != "" {
		query := gountries.New()
		c, err2 := query.FindCountryByAlpha(s.Country)
		if err2 != nil {
			return
		}
		s.CountryName = c.Name.Official
	}

	line = strings.ToLower(line)

	// Find Action
	if strings.HasPrefix(line, "invalid user") {
		s.Type = InvalidUser
		return
	}

	if strings.HasPrefix(line, "did not receive identification string") {
		s.Type = DidNotReceiveIdentificationString
		return
	}

	if strings.HasPrefix(line, "disconnected from authenticating") {
		s.Type = DisconnectedFromAuthenticating
		return
	}

	if strings.Contains(line, "session opened") {
		s.Type = SessionOpened
		return
	}

	return
}
