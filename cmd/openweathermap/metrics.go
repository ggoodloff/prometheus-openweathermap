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
	WetBulb     *prometheus.GaugeVec

	Pressure *prometheus.GaugeVec
	Humidity *prometheus.GaugeVec
	DewPoint *prometheus.GaugeVec

	UVIndex *prometheus.GaugeVec

	Clouds     *prometheus.GaugeVec
	Visibility *prometheus.GaugeVec

	WindSpeed *prometheus.GaugeVec
	WindGust  *prometheus.GaugeVec

	Weather     *prometheus.GaugeVec
	WeatherIcon *prometheus.GaugeVec

	Alerts *prometheus.GaugeVec
}

func newMetrics(reg *prometheus.Registry) *metrics {
	factory := promauto.With(reg)
	return &metrics{
		PollutionComponents: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Subsystem: "air_pollution",
			Name:      "components",
			Help:      "Air pollutant concentration in micrograms per cubic meter",
		}, []string{"station", "component"}),
		AirQualityIndex: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Subsystem: "air_pollution",
			Name:      "aqi",
			Help:      "Current air quality index (1-5)",
		}, []string{"station"}),

		Temperature: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Subsystem: "onecall",
			Name:      "temperature",
		}, []string{"station"}),
		FeelsLike: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Subsystem: "onecall",
			Name:      "feels_like",
		}, []string{"station"}),
		WetBulb: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Subsystem: "onecall",
			Name:      "wet_bulb",
		}, []string{"station"}),

		Pressure: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Subsystem: "onecall",
			Name:      "pressure",
		}, []string{"station"}),
		Humidity: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Subsystem: "onecall",
			Name:      "humidity",
		}, []string{"station"}),
		DewPoint: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Subsystem: "onecall",
			Name:      "dew_point",
		}, []string{"station"}),

		UVIndex: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Subsystem: "onecall",
			Name:      "uvi",
		}, []string{"station"}),

		Clouds: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Subsystem: "onecall",
			Name:      "clouds",
		}, []string{"station"}),
		Visibility: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Subsystem: "onecall",
			Name:      "visibility",
		}, []string{"station"}),

		WindSpeed: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Subsystem: "onecall",
			Name:      "wind_speed",
		}, []string{"station"}),
		WindGust: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Subsystem: "onecall",
			Name:      "wind_gust",
		}, []string{"station"}),

		Weather: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Subsystem: "onecall",
			Name:      "weather",
			Help:      "Current weather condition codes (https://openweathermap.org/weather-conditions)",
		}, []string{"station", "code"}),
		WeatherIcon: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Subsystem: "onecall",
			Name:      "icon",
			Help:      "Current weather condition icons (https://openweathermap.org/weather-conditions)",
		}, []string{"station", "icon"}),

		Alerts: factory.NewGaugeVec(prometheus.GaugeOpts{
			Namespace: "openweathermap",
			Name:      "alerts",
		}, []string{"station", "sender", "event"}),
	}
}
