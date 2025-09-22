package main

import (
	"encoding/json"
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gorilla/websocket"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (
	tokensTotal = prometheus.NewCounterVec(
		prometheus.CounterOpts{
			Name: "unshackled_tokens_generated_total",
			Help: "Total uncensored tokens generated since exporter start",
		}, []string{"model"})

	tokensPerSec = prometheus.NewGauge(
		prometheus.GaugeOpts{
			Name: "unshackled_tokens_per_second",
			Help: "Real-time token throughput",
		})

	secret   = getenv("COUNTER_SECRET", "unshackled")
	wsURL    = getenv("TOKEN_SOCKET", "ws://localhost:8765")
	modelTag = getenv("MODEL_TAG", "dan") // label value
)

func getenv(k, d string) string {
	if v := os.Getenv(k); v != "" {
		return v
	}
	return d
}

func main() {
	prometheus.MustRegister(tokensTotal, tokensPerSec)

	go scrapeWS()

	http.Handle("/metrics", promhttp.Handler())
	log.Fatal(http.ListenAndServe(":9100", nil))
}

func scrapeWS() {
	var mu sync.Mutex
	var lastCount uint64
	var lastTime = time.Now()

	for {
		func() {
			c, _, err := websocket.DefaultDialer.Dial(wsURL, nil)
			if err != nil {
				time.Sleep(5 * time.Second)
				return
			}
			defer c.Close()

			for {
				_, msg, err := c.ReadMessage()
				if err != nil {
					return
				}
				var m struct {
					Secret string `json:"secret"`
					Delta  uint64 `json:"delta"`
				}
				if json.Unmarshal(msg, &m) != nil || m.Secret != secret {
					continue
				}
				tokensTotal.WithLabelValues(modelTag).Add(float64(m.Delta))

				mu.Lock()
				now := time.Now()
				elapsed := now.Sub(lastTime).Seconds()
				if elapsed >= 1.0 {
					tokensPerSec.Set(float64(lastCount) / elapsed)
					lastCount = 0
					lastTime = now
				}
				lastCount += m.Delta
				mu.Unlock()
			}
		}()
	}
}
