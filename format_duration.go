package main

import (
    "time"
    "fmt"
)

func formatDuration(d time.Duration) string {
	seconds := d.Seconds()

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
