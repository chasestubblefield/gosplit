package main

import "time"

type SplitsData struct {
	Game
	Splits []Split
}

type Game struct {
	Name, Category string
}

type Split struct {
	Name         string
	PersonalBest time.Duration "pb"
	BestSegment  time.Duration "best_segment"
}
