package main

import (
	"os"

	log "github.com/sirupsen/logrus"

	sshlogexporter "github.com/DenBeke/ssh-log-exporter"
)

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}

var (
	HTTPPort = 9090
	LOGFILE  = getEnv("LOGFILE", "/var/log/auth.log")
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
