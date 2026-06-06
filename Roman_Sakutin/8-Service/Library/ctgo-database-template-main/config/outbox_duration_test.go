package config

import (
	"fmt"
	"testing"
	"time"
)

func TestParseOutboxDurationMS_NoOpOutboxWait(t *testing.T) {
	raw := fmt.Sprint(int(10_000_000 * time.Millisecond))
	got := ParseOutboxDurationMS(raw)

	want := 10_000 * time.Second
	if got != want {
		t.Fatalf("expected %v, got %v", want, got)
	}
}

func TestParseOutboxDurationMS_NormalWait(t *testing.T) {
	got := ParseOutboxDurationMS("3000")
	if got != 3*time.Second {
		t.Fatalf("expected 3s, got %v", got)
	}
}
