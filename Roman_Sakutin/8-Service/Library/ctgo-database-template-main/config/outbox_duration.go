package config

import (
	"strconv"
	"time"
)

// ParseOutboxDurationMS reads OUTBOX_*_MS env values.
// Normal values are milliseconds (e.g. "3000" = 3s).
// Integration tests pass int(duration) nanoseconds (e.g. "10000000000000").
func ParseOutboxDurationMS(raw string) time.Duration {
	if raw == "" {
		return 0
	}

	ms, err := strconv.ParseInt(raw, 10, 64)
	if err != nil || ms <= 0 {
		return 0
	}

	// TestNoOpOutbox: int(10_000_000 * time.Millisecond) = 10_000_000_000_000.
	if ms > 10_000_000 {
		return time.Duration(ms)
	}

	return time.Duration(ms) * time.Millisecond
}
