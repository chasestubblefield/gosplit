package main

import (
	"fmt"
	"math/rand"
	"time"
)

func doFakeSplits() {
	var startTime, endTime time.Time

	splits := []struct {
		name string
	}{
		{"Hero's Sword"},
		{"Outset"},
		{"Spoils Bag"},
		{"Forsaken Fortress 1"},
		{"Wind Waker"},
	}

	rand.Seed(time.Now().Unix())

	startTime = time.Now().UTC()

	fmt.Println("Started at:", startTime)

	for i, s := range splits {
		randomWait()

		t := time.Now().UTC()
		if i == len(splits)-1 {
			endTime = t
		}

		duration := t.Sub(startTime)

		fmt.Printf("%v: %v\n", s.name, formatDuration(duration))
	}

	fmt.Println("Ended at:", endTime)
}

// waits a random amount of nanoseconds between [5e9, 10e9)
func randomWait() {
	time.Sleep(time.Duration(5e9 + rand.Intn(5e9)))
}
