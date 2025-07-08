package sdf

import "fmt"
import "time"

func Example() {
	t := time.Date(2020, 10, 12, 0, 0, 0, 0, time.UTC)
	fmt.Println(Format(t))
	// Output: P0A12
}
