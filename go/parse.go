package sdf

import (
	"errors"
	"regexp"
	"strconv"
	"strings"
	"time"
)

var (
	errFormat = errors.New("bad format")
	reYear    = regexp.MustCompile(`^([D-W])(\d)([1-9ABC])([0-3]?)(\d?)`)
	reLN      = regexp.MustCompile(`^([A-Z])(\d)`)
	reValid   = regexp.MustCompile(`^([A-Z]|[A-Z]\d|[D-W]\d[1-9ABC][0-3]?\d?)`)
)

// Parse converts SDF back to milliseconds since epoch.
func Parse(s string) (int64, error) {
	s = strings.ToUpper(s)
	if !reValid.MatchString(s) {
		return 0, errFormat
	}
	if m := reYear.FindStringSubmatch(s); m != nil {
		return parseYear(m)
	}
	if m := reLN.FindStringSubmatch(s); m != nil {
		return parseLN(m)
	}
	return parseL(s)
}

func parseYear(m []string) (int64, error) {
	t, err := parseLN(m[:3])
	if err != nil {
		return 0, err
	}
	if t == infNeg || t == infPos {
		return t, nil
	}
	y, _, _ := time.UnixMilli(t).Date()
	mo := strings.IndexByte(months, m[3][0])
	if mo < 0 {
		return 0, errFormat
	}
	dd := 0
	if m[4]+m[5] != "" {
		d1 := m[4]
		d2 := m[5]
		if d1 == "" || d1 == "0" && d2 == "0" {
			return 0, errFormat
		}
		if d2 == "" {
			d2 = "0"
		}
		dd, _ = strconv.Atoi(d1 + d2)
	}
	u := time.Date(y, time.Month(mo+1), 1, 0, 0, 0, 0, time.UTC)
	if dd > 0 {
		u = u.AddDate(0, 0, dd-1)
		if u.Month() != time.Month(mo+1) {
			return 0, errFormat
		}
	}
	return u.UnixMilli(), nil
}

func parseLN(m []string) (int64, error) {
	t, err := parseL(m[1])
	if err != nil || t == infNeg || t == infPos {
		return t, err
	}
	y, _, _ := time.UnixMilli(t).Date()
	d := int(m[2][0] - '0')
	switch m[1] {
	case "B":
		if d > 7 {
			return 0, errFormat
		}
		y += d * 100
	case "Y":
		if d < 2 {
			return 0, errFormat
		}
		y = 2000 + d*100
	case "C", "X":
		y += d * 10
	default:
		y += d
	}
	return time.Date(y, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), nil
}

const (
	infNeg = int64(-1 << 63)
	infPos = int64(^uint64(0) >> 1)
)

func parseL(s string) (int64, error) {
	switch s {
	case "A":
		return infNeg, nil
	case "B":
		return time.Date(1000, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), nil
	case "C":
		return time.Date(1800, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), nil
	case "X":
		return time.Date(2100, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), nil
	case "Y":
		return time.Date(2200, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), nil
	case "Z":
		return infPos, nil
	}
	y := int(s[0]-'D')*10 + 1900
	return time.Date(y, 1, 1, 0, 0, 0, 0, time.UTC).UnixMilli(), nil
}
