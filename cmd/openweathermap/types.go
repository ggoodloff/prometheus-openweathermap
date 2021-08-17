package main

import (
	"context"
	"time"
)

type config struct {
	Address string `mapstructure:"address"`

	API openweathermap `mapstructure:"api"`

	Stations []station `mapstructure:"stations"`
}

type openweathermap struct {
	Key     string `mapstructure:"key"`
	BaseURL string `mapstructure:"base_url"`

	MaxCallsPerMonth uint          `mapstructure:"max_calls_per_month"`
	MinPollRate      time.Duration `mapstructure:"min_poll_rate"`
	Backoff          backoff       `mapstructure:"backoff"`
}

type station struct {
	Name      string  `mapstructure:"name"`
	Latitude  float64 `mapstructure:"latitude"`
	Longitude float64 `mapstructure:"longitude"`
	Metrics   struct {
		Pollution bool `mapstructure:"pollution"`
	}
}

type coord struct {
	Latitude  string `mapstructure:"lat"`
	Longitude string `mapstructure:"lon"`
}

type collectorFunc func(context.Context) error
