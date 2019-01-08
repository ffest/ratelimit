package ratelimit

import (
	"testing"
	"time"
)

func TestRateLimiter_Wait_ShouldBlock(t *testing.T) {
	startTime := time.Now()
	limit := 100
	interval := time.Second
	limiter := New(limit, interval)

	for i := 0; i < limit+1; i++ {
		limiter.Wait()
	}

	if time.Now().Sub(startTime) < interval {
		t.Error("The limiter didn't block enough!")
	}
}

func TestRateLimiter_Wait_ShouldNotBlock(t *testing.T) {
	startTime := time.Now()
	limit := 10
	interval := time.Second
	limiter := New(limit, interval)

	for i := 0; i < limit-1; i++ {
		limiter.Wait()
	}

	if time.Now().Sub(startTime) >= interval {
		t.Error("The limiter blocked more than need!")
	}
}
