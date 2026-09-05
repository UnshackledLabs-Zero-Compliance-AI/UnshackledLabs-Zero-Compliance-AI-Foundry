package main

import (
    "net/http"
    "github.com/prometheus/client_golang/prometheus"
    "github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
    requestsTotal = prometheus.NewCounterVec(
        prometheus.CounterOpts{
            Name: "inference_requests_total",
            Help: "Total inference requests",
        },
        []string{"status"},
    )
    activeConns = prometheus.NewGauge(
        prometheus.GaugeOpts{
            Name: "inference_active_connections",
            Help: "Active WebSocket connections",
        },
    )
)

func init() {
    prometheus.MustRegister(requestsTotal, activeConns)
}

func main() {
    // Simulate metrics update (in real world, you'd scrape from relay/inference)
    go func() {
        // dummy updater for demo
    }()
    
    http.Handle("/metrics", promhttp.Handler())
    http.ListenAndServe(":9090", nil)
}