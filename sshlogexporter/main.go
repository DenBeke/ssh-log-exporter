package main

import (
	log "github.com/sirupsen/logrus"

	sshlogexporter "github.com/DenBeke/ssh-log-exporter"
)

var (
	HTTPPort = 9090
	LOGFILE  = "/var/log/auth.log"
)

func main() {

	go func() {
		sshlogexporter.RunPrometheusServer(HTTPPort)
	}()

	err := sshlogexporter.TailAuthLog(LOGFILE)
	if err != nil {
		log.Fatalf("couldn't tail file: %v", err)
	}

}
