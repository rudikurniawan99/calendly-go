package main

import (
	"fmt"
	"time"
)

func main() {
	// now := time.Now().UTC()
	now := time.Date(2022, time.August, 20, 0, 0, 0, 0, time.Local)
	duration := time.Duration(30) * time.Minute

	now = now.Add(duration)

	// fmt.Println(now.Add(duration))
	// now = now.Add(duration)
	fmt.Println(now)
}
