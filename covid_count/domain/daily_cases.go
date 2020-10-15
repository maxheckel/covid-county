package domain

import "time"

type DailyCases struct {
	County string    `json:"county"`
	Date   time.Time `json:"date"`
	Count  int       `json:"count"`
}
