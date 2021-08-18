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

type onecall struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	Timezone       string `json:"timezone"`
	TimezoneOffset int    `json:"timezone_offset"`

	Current struct {
		Timestamp int64 `json:"dt"`

		Sunrise int64 `json:"sunrise"`
		Sunset  int64 `json:"sunset"`

		Temperature float64 `json:"temp"`
		FeelsLike   float64 `json:"feels_like"`
		Pressure    float64 `json:"pressure"`
		Humidity    float64 `json:"humidity"`
		DewPoint    float64 `json:"dew_point"`

		UVIndex float64 `json:"uvi"`

		Clouds     float64 `json:"clouds"`
		Visibility float64 `json:"visibility"`

		WindSpeed     float64 `json:"wind_speed"`
		WindGust      float64 `json:"wind_gust"`
		WindDirection float64 `json:"wind_deg"`

		Weather []struct {
			ID          int    `json:"id"`
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
	} `json:"current"`

	Minutely []struct {
	} `json:"minutely"`

	Hourly []struct {
	} `json:"hourly"`

	Daily []struct {
	} `json:"daily"`
}

func (env *environment) collectWeather(s station) collectorFunc {
	endpoint := env.BaseURL
	endpoint.Path = path.Join(endpoint.Path, "onecall")

	labels := prometheus.Labels{
		"station": s.Name,
	}

	metricWeather := env.Metrics.Weather.MustCurryWith(labels)
	metricWeatherIcon := env.Metrics.WeatherIcon.MustCurryWith(labels)

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

		var data onecall
		if err = json.NewDecoder(res.Body).Decode(&data); err != nil {
			return err
		}

		env.Metrics.Temperature.With(labels).Set(data.Current.Temperature)
		env.Metrics.FeelsLike.With(labels).Set(data.Current.FeelsLike)

		env.Metrics.Pressure.With(labels).Set(data.Current.Pressure)
		env.Metrics.Humidity.With(labels).Set(data.Current.Humidity)
		env.Metrics.DewPoint.With(labels).Set(data.Current.DewPoint)

		env.Metrics.UVIndex.With(labels).Set(data.Current.UVIndex)

		env.Metrics.Clouds.With(labels).Set(data.Current.Clouds)
		env.Metrics.Visibility.With(labels).Set(data.Current.Visibility)

		env.Metrics.WindSpeed.With(labels).Set(data.Current.WindSpeed)
		env.Metrics.WindGust.With(labels).Set(data.Current.WindGust)

		metricWeather.Reset()
		metricWeatherIcon.Reset()
		for _, w := range data.Current.Weather {
			metricWeather.With(prometheus.Labels{
				"code": fmt.Sprint(w.ID),
			}).Set(1)

			metricWeatherIcon.With(prometheus.Labels{
				"icon": fmt.Sprint(w.Icon),
			}).Set(1)
		}

		if env.Units == "metric" {
			env.Metrics.WetBulb.With(labels).Set(CalculateWetBulbTemperature(
				data.Current.Temperature,
				data.Current.Humidity,
			))
		}

		return nil
	}
}
