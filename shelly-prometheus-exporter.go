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

	"shelly-prometheus-exporter/shelly"
	"shelly-prometheus-exporter/shelly_detect"
	"shelly-prometheus-exporter/shelly_v1"
	"shelly-prometheus-exporter/shelly_v2"
)

var addr = flag.String("listen-address", ":54901", "The address to listen on for HTTP requests.")

const userAgent = "shelly-prometheus-exporter"

func probeHandler(w http.ResponseWriter, req *http.Request) {
	reg := prometheus.NewPedanticRegistry()

	m := shelly.NewMetrics(reg)

	ctx, cancel := context.WithCancel(req.Context())
	defer cancel()
	req = req.WithContext(ctx)

	target := req.URL.Query().Get("target")
	if target == "" {
		http.Error(w, "'target' parameter is missing", http.StatusBadRequest)
		return
	}
	username := req.URL.Query().Get("username")
	password := req.URL.Query().Get("password")

	shellyType, err := shelly_detect.DetectVersion(target)
	if err != nil {
		http.Error(w, "unable to detect the shelly version", http.StatusInternalServerError)
		return
	}

	var s shelly.Device
	switch shellyType {
	case shelly_detect.ShellyGeneration1:
		s = shelly_v1.New(target, userAgent, username, password)
	case shelly_detect.ShellyGeneration2:
		s = shelly_v2.New(target, userAgent, password)
	default:
		http.Error(w, fmt.Sprintf("unkown shelly generation '%d'", shellyType), http.StatusBadRequest)
		return
	}

	err = s.FetchStatus()
	if err != nil {
		http.Error(w, fmt.Sprintf("error while fetching status: %s", err), http.StatusServiceUnavailable)
		return
	}
	s.FillMetrics(m)

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
