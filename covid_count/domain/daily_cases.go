package domain

import "time"

type InstanceType string

const (
	Death           InstanceType = "Death"
	Case            InstanceType = "Case"
	Hospitalization InstanceType = "Hospitalization"
)

type DailyInstances struct {
	Type   InstanceType `json:"type"`
	County string       `json:"county"`
	Date   time.Time    `json:"date"`
	Count  int          `json:"count"`
}
