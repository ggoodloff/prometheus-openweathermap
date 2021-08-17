package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type metrics struct {
	AirPollution    *prometheus.GaugeVec
	AirQualityIndex *prometheus.GaugeVec
}

func newMetrics(reg *prometheus.Registry) *metrics {
	factory := promauto.With(reg)
	return &metrics{
		AirPollution: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Name:      "air_pollution",
		}, []string{"station", "component"}),
		AirQualityIndex: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Name:      "air_quality_index",
		}, []string{"station"}),
	}
}
