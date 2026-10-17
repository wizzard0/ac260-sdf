package sdf

import (
	"math"
	"math/rand"
	"testing"
	"time"
)

// why: SDF promises that codes sort like the dates they encode, in every range
// given: random dates from year -5000 to 5000 and random forward offsets
// when: x and x + y are both formatted and parsed back
// then: neither the code nor its decoded start ever goes backwards
func TestOrderPreserved(t *testing.T) {
	lo := time.Date(-5000, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli()
	hi := time.Date(5000, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli()
	r := rand.New(rand.NewSource(260)) // seeded, so a failure always reproduces
	for i := 0; i < 20000; i++ {
		x := lo + r.Int63n(hi-lo)
		y := int64(math.Pow(r.Float64(), 4) * float64(hi-x))
		a, b := Format(time.UnixMilli(x)), Format(time.UnixMilli(x+y))
		if a > b {
			t.Fatalf("code order: %s=%s > +%d=%s", time.UnixMilli(x).UTC(), a, y, b)
		}
		pa, errA := Parse(a)
		pb, errB := Parse(b)
		if errA != nil || errB != nil || pa > pb {
			t.Fatalf("parse order: %s=%d (%v) > %s=%d (%v)", a, pa, errA, b, pb, errB)
		}
	}
}
