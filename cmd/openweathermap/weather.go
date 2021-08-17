package main

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"path"

	"github.com/prometheus/client_golang/prometheus"
)

type weather struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	Timezone  int    `json:"timezone"`
	Coord     coord  `json:"coord"`
	Timestamp uint   `json:"dt"`

	Sys struct {
		Sunrise int `json:"sunrise"`
		Sunset  int `json:"suset"`
	} `json:"sys"`

	Weather []struct {
		ID          int    `json:"id"`
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`

	Main struct {
		Temperature float64 `json:"temp"`
		FeelsLike   float64 `json:"feels_like"`

		MinTemperature float64 `json:"temp_min"`
		MaxTemperature float64 `json:"temp_max"`

		Humidity float64 `json:"humidity"`

		Pressure            float64 `json:"pressure"`
		SeaLevelPressure    float64 `json:"sea_level"`
		GroundLevelPressure float64 `json:"grnd_level"`
	} `json:"main"`

	Wind struct {
		Speed     float64 `json:"speed"`
		Gust      float64 `json:"gust"`
		Direction float64 `json:"deg"`
	} `json:"wind"`

	Clouds struct {
		Cloudiness float64 `json:"all"`
	} `json:"clouds"`

	Rain struct {
	} `json:"rain"`

	Snow struct {
	} `json:"snow"`
}

func (env *environment) collectWeather(s station) collectorFunc {
	endpoint := env.BaseURL
	endpoint.Path = path.Join(endpoint.Path, "weather")

	labels := prometheus.Labels{
		"station": s.Name,
	}

	return func(ctx context.Context) error {
		url := endpoint

		q := url.Query()
		q.Add("lat", fmt.Sprintf("%.6f", s.Latitude))
		q.Add("lon", fmt.Sprintf("%.6f", s.Longitude))
		q.Add("units", env.Units)
		url.RawQuery = q.Encode()

		log.Printf("Collecting weather for station %s", s.Name)
		res, err := http.Get(url.String())
		if err != nil {
			return err
		}
		defer res.Body.Close()

		if res.StatusCode != http.StatusOK {
			return fmt.Errorf(
				"Failed to get weather for station %s: %d %s",
				s.Name, res.StatusCode, http.StatusText(res.StatusCode),
			)
		}

		var data weather
		if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
			return err
		}

		env.Metrics.Temperature.With(labels).Set(data.Main.Temperature)
		env.Metrics.FeelsLike.With(labels).Set(data.Main.FeelsLike)
		env.Metrics.MinTemperature.With(labels).Set(data.Main.MinTemperature)
		env.Metrics.MaxTemperature.With(labels).Set(data.Main.MaxTemperature)

		env.Metrics.Humidity.With(labels).Set(data.Main.Humidity)
		env.Metrics.Pressure.With(labels).Set(data.Main.Pressure)

		env.Metrics.WindSpeed.With(labels).Set(data.Wind.Speed)
		env.Metrics.WindGust.With(labels).Set(data.Wind.Gust)

		env.Metrics.Cloudiness.With(labels).Set(data.Clouds.Cloudiness)

		return nil
	}
}
