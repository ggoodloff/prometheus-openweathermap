package main

import (
	"github.com/prometheus/client_golang/prometheus"
)

type metrics struct {
}

func newMetrics(reg *prometheus.Registry) *metrics {
	return &metrics{
	}
}
