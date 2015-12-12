package main

import (
	"fmt"
	"time"
)

type splitsData struct {
	Game game
	Splits []split
}

type game struct {
	Name, Category string
}

type split struct {
	Name string
	PersonalBest time.Duration
	BestSegment time.Duration
}
