package main

import (
	"log"
	"time"
)

func main() {

	params := struct {
		y, m int
		w    time.Weekday
	}{
		y: 2022,
		m: 8,
		w: time.Monday,
	}

	star := time.Date(params.y, time.Month(params.m), 1, 0, 0, 0, 0, time.Local)
	end := time.Date(params.y, time.Month(params.m+1), 0, 0, 0, 0, 0, time.Local)

	starDay := star.Day()
	endDay := end.Day()

	var availableDay []time.Time

	for i := starDay; i <= endDay; i++ {
		cur := time.Date(params.y, time.Month(params.m), i, 0, 0, 0, 0, time.Local)

		if cur.Weekday() == params.w {

			availableDay = append(availableDay, cur)
		}

	}

	log.Println(availableDay)

}
