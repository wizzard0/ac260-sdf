package sdf

import (
	"fmt"
	"strconv"
	"time"
)

const months = "123456789ABC"

// Format returns SDF string for a time.
func Format(t time.Time) string {
	y := t.UTC().Year()
	return rangeType(y)(t.UTC())
}

func rangeType(y int) func(time.Time) string {
	switch {
	case y < 1000 || y >= 3000:
		return formatInf
	case y < 1800 || y >= 2200:
		return formatCen
	case y < 1900 || y >= 2100:
		return formatDec
	}
	return formatYear
}

func formatYear(t time.Time) string {
	y := t.Year()
	d := int(y/10) - 190
	l := string(rune('D' + d))
	m := months[t.Month()-1 : t.Month()]
	return l + strconv.Itoa(y%10) + m + two(t.Day())
}

func formatDec(t time.Time) string {
	y := t.Year()
	d := y % 100 / 10
	pre := 'C'
	if y >= 2100 {
		pre = 'X'
	}
	return string(pre) + strconv.Itoa(d)
}

func formatCen(t time.Time) string {
	y := t.Year()
	c := y % 1000 / 100
	pre := 'B'
	if y >= 2200 {
		pre = 'Y'
	}
	return string(pre) + strconv.Itoa(c)
}

func formatInf(t time.Time) string {
	if t.Before(time.Unix(0, 0)) {
		return "A0"
	}
	return "Z9"
}

func two(d int) string { return fmt.Sprintf("%02d", d) }
