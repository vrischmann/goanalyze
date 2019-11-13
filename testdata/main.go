package main

import "github.com/prometheus/client_golang/prometheus"

var (
	histogram = prometheus.NewHistogram(prometheus.HistogramOpts{
		Name:      "histogram",
		Help:      "help",
		Namespace: "prom",
	})
)

func main() {
}
