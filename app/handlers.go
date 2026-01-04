package main

import (
	"fmt"
	"net/http"
	"time"
)

func healthHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintln(w, "OK")
}

func loadHandler(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	for time.Since(start) < 2*time.Second {
		// busy loop to generate CPU load
	}
	fmt.Fprintln(w, "Load generated")
}
