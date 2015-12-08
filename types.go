package main

import (
	"fmt"
	"time"
)

type duration time.Duration

func (d duration) String() string {
	seconds := time.Duration(d).Seconds()

	minutes := int(seconds / 60)
	seconds -= float64(minutes * 60)

	hours := int(minutes / 60)
	minutes -= hours * 60

	if hours > 0 {
		return fmt.Sprintf("%d:%02d:%05.2f", hours, minutes, seconds)
	} else {
		return fmt.Sprintf("%d:%05.2f", minutes, seconds)
	}
}

type split struct {
	name string
}

type run struct {
	start, end time.Time
}
