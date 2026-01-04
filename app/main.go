package main

import (
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	http.HandleFunc("/health", healthHandler)
	http.HandleFunc("/load", loadHandler)
	http.Handle("/metrics", promhttp.Handler())

	log.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
