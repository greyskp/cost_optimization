package optimization

import "time"

// Stats is a lightweight, backend-agnostic payload that callers can export to logs/metrics/traces however they want.
type Stats struct {
	N             int
	SelectedCount int
	LeftToFill    int
	Replacements  int
	Duration      time.Duration
}

// Observer is an optional hook. Implementations should be fast and non-blocking
type Observer interface {
	Observe(stats Stats)
}

// NoOpObserver is the default when no observer is provided.
type NoOpObserver struct{}

func (NoOpObserver) Observe(Stats) {}