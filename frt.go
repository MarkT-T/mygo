package main

import (
	"fmt"
	"time"
)

func main() {
	frtdate := time.Date(1789, 6, 14, 0, 0, 0, 0, time.UTC)
	now := time.Now()
	diff := now.Sub(frtdate)

	fmt.Println("The Great French Revolution began", int64(diff.Hours()/24/365), "years,", int64(diff.Hours()/24/30), "month or", int64(diff.Hours()/24), "days ago")
}
