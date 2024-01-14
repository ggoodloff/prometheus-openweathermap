package main

import (
	"net/url"

	"github.com/prometheus/client_golang/prometheus"
)

type environment struct {
	Registry *prometheus.Registry
	Metrics  *metrics

	BaseURL url.URL
	Units   string
}

func newEnvironment(cfg *config) (*environment, error) {
	r := prometheus.NewRegistry()

	baseURL, err := url.Parse(cfg.API.BaseURL)
	if err != nil {
		return nil, err
	}

	q := baseURL.Query()
	q.Add("appid", cfg.API.Key)
	baseURL.RawQuery = q.Encode()

	return &environment{
		Registry: r,
		Metrics:  newMetrics(r),
		BaseURL:  *baseURL,
		Units:    cfg.API.Units,
	}, nil
}
