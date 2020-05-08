package sshlogexporter

import (
	"github.com/hpcloud/tail"

	log "github.com/sirupsen/logrus"
)

func TailAuthLog(logFile string) error {

	t, err := tail.TailFile(logFile, tail.Config{Follow: true})

	if err != nil {
		return err
	}

	for line := range t.Lines {
		s, err := ExtractSSHFromLine(line.Text)
		if err != nil {
			log.Warnf("something went wrong while parsing the line: %v", err)
		}
		if s == nil {
			continue
		}

		if s.IsAttackAttempt() {
			sshAttempts.WithLabelValues(s.IP, s.Country).Add(1)
		}

		log.WithField("ip", s.IP).WithField("country", s.Country).WithField("username", s.Username).WithField("attack_attempt", s.IsAttackAttempt()).Printf(s.Log)
	}

	return nil

}
