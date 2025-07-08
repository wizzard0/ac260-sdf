package sdf

import (
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	ts := time.Date(2020, 10, 12, 0, 0, 0, 0, time.UTC)
	got := Format(ts)
	if got != "P0A12" {
		t.Fatalf("want %q got %q", "P0A12", got)
	}
}
