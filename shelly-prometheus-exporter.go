package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/collectors"
	"github.com/prometheus/client_golang/prometheus/promhttp"

	"shelly-prometheus-exporter/shelly_metrics"
	"shelly-prometheus-exporter/shelly_v1"
)

var addr = flag.String("listen-address", ":54901", "The address to listen on for HTTP requests.")

func probeHandler(w http.ResponseWriter, req *http.Request) {
	reg := prometheus.NewPedanticRegistry()

	m := shelly_metrics.NewMetrics(reg)

	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()
	req = req.WithContext(ctx)

	target := req.URL.Query().Get("target")
	if target == "" {
		http.Error(w, "'target' parameter is missing", http.StatusBadRequest)
		return
	}
	shellyType := req.URL.Query().Get("shelly_type")
	if target == "" {
		http.Error(w, "'shelly_type' parameter is missing", http.StatusBadRequest)
		return
	}

	if shellyType == "v1" {
		data, err := shelly_v1.FetchStatus(target)
		if err != nil {
			http.Error(w, fmt.Sprintf("error while fetching status: %s", err), http.StatusServiceUnavailable)
			return
		}
		shelly_v1.ParseMetrics(data, m)
	} else if shellyType == "v2" {
		// TODO v2
	} else {
		http.Error(w, fmt.Sprintf("unkown shelly_type '%s'", shellyType), http.StatusBadRequest)
	}

	h := promhttp.HandlerFor(reg, promhttp.HandlerOpts{EnableOpenMetrics: true})
	h.ServeHTTP(w, req)
}

func main() {
	flag.Parse()

	registry := prometheus.NewRegistry()

	registry.MustRegister(
		collectors.NewGoCollector(),
		collectors.NewBuildInfoCollector(),
	)

	http.Handle("/metrics",
		promhttp.HandlerFor(registry,
			promhttp.HandlerOpts{
				EnableOpenMetrics: true,
				Registry:          registry,
			}))
	http.HandleFunc("/probe", probeHandler)

	log.Printf("Now listening on '%s'…\n", *addr)
	log.Fatalln(http.ListenAndServe(*addr, nil))
}
