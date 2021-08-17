package main

import (
	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promauto"
)

type metrics struct {
	PollutionComponents *prometheus.GaugeVec
	AirQualityIndex     *prometheus.GaugeVec

	Temperature *prometheus.GaugeVec
	FeelsLike   *prometheus.GaugeVec

	MinTemperature *prometheus.GaugeVec
	MaxTemperature *prometheus.GaugeVec

	Humidity *prometheus.GaugeVec
	Pressure *prometheus.GaugeVec

	WindSpeed *prometheus.GaugeVec
	WindGust  *prometheus.GaugeVec

	Cloudiness *prometheus.GaugeVec
}

func newMetrics(reg *prometheus.Registry) *metrics {
	factory := promauto.With(reg)
	return &metrics{
		PollutionComponents: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Subsystem: "air_pollution",
			Name:      "components",
		}, []string{"station", "component"}),
		AirQualityIndex: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Subsystem: "air_pollution",
			Name:      "aqi",
		}, []string{"station"}),

		Temperature: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Name:      "temperature",
		}, []string{"station"}),
		FeelsLike: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Name:      "feels_like",
		}, []string{"station"}),
		MinTemperature: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Name:      "min_temperature",
		}, []string{"station"}),
		MaxTemperature: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Name:      "max_temperature",
		}, []string{"station"}),

		Humidity: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Name:      "humidity",
		}, []string{"station"}),
		Pressure: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Name:      "pressure",
		}, []string{"station"}),

		WindSpeed: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Name:      "wind_speed",
		}, []string{"station"}),
		WindGust: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Name:      "wind_gust",
		}, []string{"station"}),

		Cloudiness: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Name:      "cloudiness",
		}, []string{"station"}),
	}
}
