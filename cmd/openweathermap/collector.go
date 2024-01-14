package main

import (
	"context"
	"log"
	"time"
)

type collector struct {
	Collect collectorFunc
	Rate    time.Duration
	Backoff backoff
}

func (c *collector) Run(ctx context.Context) {
	for {
		var d time.Duration

		err := c.Collect(ctx)
		if err != nil {
			d = c.Backoff.Current()
			c.Backoff.Backoff()

			log.Printf("Collector failure: %v\n", err)
		} else {
			c.Backoff.Reset()
			d = time.Until(time.Now().Add(c.Rate))
		}

		t := time.NewTimer(d)
		select {
		case <-ctx.Done():
			return
		case <-t.C:
			continue
		}
	}
}
