package main

import "time"

type backoff struct {
	BaseDelay  time.Duration `mapstructure:"base_delay"`
	MaxDelay   time.Duration `mapstructure:"max_delay"`
	Multiplier float64       `mapstructure:"multiplier"`

	value time.Duration
}

func (b *backoff) Current() time.Duration {
	return b.value
}

func (b *backoff) Reset() time.Duration {
	b.value = b.BaseDelay
	return b.value
}

func (b *backoff) Backoff() {
	b.value = time.Duration(b.Multiplier * float64(b.value))
	if b.value > b.MaxDelay {
		b.value = b.MaxDelay
	}
}
