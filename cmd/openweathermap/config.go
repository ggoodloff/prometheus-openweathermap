package main

import (
	"fmt"
	"strings"
	"time"

	"github.com/spf13/viper"
)

func loadConfig() (*config, error) {
	v := viper.New()

	v.SetConfigName("openweathermap")
	v.SetConfigType("yml")

	v.AddConfigPath("/etc")
	v.AddConfigPath("$HOME")
	v.AddConfigPath(".")

	err := v.ReadInConfig()
	if err != nil {
		panic(fmt.Errorf("Fatal error config file: %s", err))
	}

	v.SetEnvPrefix("owm")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))
	v.AutomaticEnv()

	populateDefaultConfig(v)

	var c config
	err = v.Unmarshal(&c)
	if err != nil {
		return nil, err
	}

	return &c, nil
}

func populateDefaultConfig(v *viper.Viper) {
	v.SetDefault("address", ":80")
	v.SetDefault("api", map[string]interface{}{
		"base_url":                "http://api.openweathermap.org/data/2.5/",
		"max_calls_per_month": 10000,
		"min_poll_rate":           time.Minute,
		"backoff": map[string]interface{}{
			"base_delay": 10 * time.Second,
			"max_delay":  10 * time.Minute,
			"multiplier": 2,
		},
	})
	v.SetDefault("metrics", map[string]interface{}{
		"pollution": true,
	})
}
