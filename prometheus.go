package sshlogexporter

import (
	"fmt"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	sshAttempts = promauto.NewCounterVec(
		prometheus.CounterOpts{
			Name: "ssh_attack_attempts_total",
			Help: "How many SSH actions processed, partitioned by ...",
		},
		[]string{"ip", "country"},
	)
)

func RunPrometheusServer(port int) error {
	http.Handle("/metrics", promhttp.Handler())
	return http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
}
