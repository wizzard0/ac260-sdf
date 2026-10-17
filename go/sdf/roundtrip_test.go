package sdf

import (
	"testing"
	"time"
)

// why: decade and century codes must decode to the period they were made from
func TestRoundTripDecadesCenturies(t *testing.T) {
	cases := []struct {
		year  int
		code  string
		start int
	}{
		{1000, "B0", 1000}, {1250, "B2", 1200}, {1799, "B7", 1700},
		{1800, "C0", 1800}, {1850, "C5", 1850}, {1899, "C9", 1890},
		{2100, "X0", 2100}, {2150, "X5", 2150}, {2199, "X9", 2190},
		{2200, "Y2", 2200}, {2750, "Y7", 2700}, {2999, "Y9", 2900},
	}
	for _, c := range cases {
		got := Format(time.Date(c.year, 6, 15, 0, 0, 0, 0, time.UTC))
		if got != c.code {
			t.Errorf("%d: want %q got %q", c.year, c.code, got)
			continue
		}
		ms, err := Parse(got)
		if err != nil {
			t.Errorf("%s: %v", got, err)
			continue
		}
		if y := time.UnixMilli(ms).UTC().Year(); y != c.start {
			t.Errorf("%s: want year %d got %d", got, c.start, y)
		}
	}
}

// why: the formatter emits A0 and Z9, so the parser must accept them
func TestParseInfinities(t *testing.T) {
	for code, want := range map[string]int64{"A0": infNeg, "Z9": infPos, "a": infNeg, "z": infPos} {
		got, err := Parse(code)
		if err != nil || got != want {
			t.Errorf("%s: want %d got %d (%v)", code, want, got, err)
		}
	}
}
