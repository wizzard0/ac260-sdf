package main

import (
	"fmt"
	"time"

	sdf "github.com/wizzard0/ac260-sdf/go"
)

func main() {
	t := time.Date(2020, 10, 12, 0, 0, 0, 0, time.UTC)
	fmt.Println(sdf.Format(t))
	ts, _ := sdf.Parse("P0A12")
	fmt.Println(time.UnixMilli(ts).UTC())
}
