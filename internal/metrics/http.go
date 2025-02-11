package metrics

import (
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"log"
	"net/http"
)

func StartHTTPServer(addr string) {
	http.Handle("/metrics", promhttp.Handler())
	log.Printf("Metrics HTTP server listening on %s/metrics", addr)
	if err := http.ListenAndServe(addr, nil); err != nil {
		log.Fatalf("Metrics HTTP server error: %v", err)
	}
}
