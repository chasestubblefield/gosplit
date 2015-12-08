package main

import (
	"fmt"
	"math/rand"
	"time"
)

var splits = []split{
	{"Hero's Sword"},
	{"Outset"},
	{"Spoils Bag"},
	{"Forsaken Fortress 1"},
	{"Wind Waker"},
}

var startTime, endTime time.Time

func main() {
	rand.Seed(time.Now().Unix())

	startTime = time.Now().UTC()

	fmt.Println("Started at:", startTime)

	for i, s := range splits {
		randomWait()

		t := time.Now().UTC()
		if i == len(splits)-1 {
			endTime = t
		}

		dur := duration(t.Sub(startTime))

		fmt.Printf("%v: %v\n", s.name, dur)
	}

	fmt.Println("Ended at:", endTime)
}

// waits a random amount of nanoseconds between [5e9, 10e9)
func randomWait() {
	time.Sleep(time.Duration(5e9 + rand.Intn(5e9)))
}
