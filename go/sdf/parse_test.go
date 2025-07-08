package sdf

import "testing"
import "time"

func TestParse(t *testing.T) {
	ts, err := Parse("N3201")
	if err != nil {
		t.Fatal(err)
	}
	want := time.Date(2003, 2, 1, 0, 0, 0, 0, time.UTC).UnixMilli()
	if ts != want {
		t.Fatalf("want %d got %d", want, ts)
	}
}
