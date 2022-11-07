package models

import "time"

type Transaction struct {
	Amount    float64   `json:"amount"`
	Timestamp time.Time `json:"timestamp"`
}

type Status struct {
	Sum   float64
	Avg   float64
	Max   float64
	Min   float64
	Count int
}
