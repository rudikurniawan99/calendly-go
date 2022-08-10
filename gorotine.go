package main

import (
	"time"
)

func GorotineDoubleIt(channel chan any, x int) {

	identity := struct {
		name  string
		age   int
		birth int
	}{
		name:  "rudi kurniawan",
		age:   x,
		birth: time.Now().Year() - x,
	}

	channel <- identity

	close(channel)
}
