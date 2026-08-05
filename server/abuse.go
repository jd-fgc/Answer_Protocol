package main

import (
	"sync"
	"time"
)

type rateTracker struct {
	mu   sync.Mutex
	hits map[string][]time.Time
}

func newRateTracker() *rateTracker {
	return &rateTracker{hits: make(map[string][]time.Time)}
}

func (r *rateTracker) hit(key string, window time.Duration, max int) bool {
	now := time.Now()
	r.mu.Lock()
	defer r.mu.Unlock()

	cutoff := now.Add(-window)
	kept := r.hits[key][:0]
	for _, t := range r.hits[key] {
		if t.After(cutoff) {
			kept = append(kept, t)
		}
	}
	kept = append(kept, now)
	r.hits[key] = kept

	return len(kept) > max
}

const (
	commandWindow       = 2 * time.Second
	commandMaxPerWindow = 15

	connectionWindow       = 5 * time.Second
	connectionMaxPerWindow = 5
)

var commandTracker = newRateTracker()
var connectionTracker = newRateTracker()

func checkAbuse(from string) bool {
	return commandTracker.hit(from, commandWindow, commandMaxPerWindow)
}

func checkConnectionAbuse(ip string) bool {
	return connectionTracker.hit(ip, connectionWindow, connectionMaxPerWindow)
}
