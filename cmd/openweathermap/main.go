package main

import (
	"context"
	"log"
	"net/http"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	ctx, shutdown := signal.NotifyContext(
		context.Background(),
		syscall.SIGINT,
		syscall.SIGTERM,
	)
	defer shutdown()

	cfg, err := loadConfig()
	if err != nil {
		log.Fatalf("Failed to load configuration: %v\n", err)
		return
	}

	cfg.API.Backoff.Reset()
	env, err := newEnvironment(cfg)
	if err != nil {
		log.Fatalf("Failed to create environment: %v\n", err)
	}

	var collectors []*collector

	log.Printf("Max API Calls Per Month: %d\n", cfg.API.MaxCallsPerMonth)
	apiCallRate := 31 * 24 * time.Hour / time.Duration(cfg.API.MaxCallsPerMonth)

	for _, s := range cfg.Stations {
		if s.Metrics.Weather {
			collectors = append(collectors, &collector{
				Collect: env.collectWeather(s),
				Rate:    apiCallRate,
				Backoff: cfg.API.Backoff,
			})
		}
		if s.Metrics.Pollution {
			collectors = append(collectors, &collector{
				Collect: env.collectPollution(s),
				Rate:    apiCallRate,
				Backoff: cfg.API.Backoff,
			})
		}
	}

	var wg sync.WaitGroup
	for _, c := range collectors {
		wg.Add(1)

		c.Rate *= time.Duration(len(collectors))
		if c.Rate < cfg.API.MinPollRate {
			c.Rate = cfg.API.MinPollRate
		}

		go func(c *collector) {
			defer wg.Done()
			c.Run(ctx)
		}(c)
	}

	mux := http.NewServeMux()
	mux.Handle("/metrics", promhttp.InstrumentMetricHandler(
		env.Registry, promhttp.HandlerFor(env.Registry, promhttp.HandlerOpts{}),
	))

	server := http.Server{
		Addr:    cfg.Address,
		Handler: mux,
	}

	wg.Add(1)
	go func() {
		defer wg.Done()
		log.Printf("Metrics server started at %s\n", cfg.Address)
		err := server.ListenAndServe()
		if err != nil {
			if err == http.ErrServerClosed {
				return
			}
			log.Printf("Failed to start metrics server: %v\n", err)
			defer shutdown()
		}
	}()

	<-ctx.Done()

	ctx, halt := context.WithTimeout(context.Background(), 10*time.Second)
	defer halt()

	err = server.Shutdown(ctx)
	if err != nil {
		log.Printf("Metrics server failed to shutdown cleanly: %v\n", err)
	}
	log.Printf("Metrics server stopped")

	wg.Wait()
	log.Printf("All collectors stopped")
}
