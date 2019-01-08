package ratelimit

import (
	"sync"
	"time"
)

// A RateLimiter is used to limits the rate of some process.
// It applies smoothing conception of warm up.
type RateLimiter struct {
	sync.Mutex
	available      int64
	ticker         *time.Ticker
	tickerDuration time.Duration
	lastRequest    time.Time
}

// New creates a new rate limiter based on limit and time interval.
// Calculates the maximum possible frequency of requests.
func New(limit int, interval time.Duration) *RateLimiter {
	tickerDuration := time.Duration(interval.Nanoseconds() / int64(limit))
	rl := &RateLimiter{
		available: 1,
		ticker:    time.NewTicker(tickerDuration),
	}
	go rl.Ticker()
	return rl
}

// Wait blocks for a while if the rate limit has been reached.
func (l *RateLimiter) Wait() {
	for {
		l.Lock()
		if l.available == 0 {
			l.Unlock()
			time.Sleep(l.tickerDuration)
			continue
		}
		l.available--
		l.Unlock()
		return
	}
}

// Ticker increments available counter per one each tick
func (l *RateLimiter) Ticker() {
	for range l.ticker.C {
		l.Lock()
		l.available++
		l.Unlock()
	}
}
