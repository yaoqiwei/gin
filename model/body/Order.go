package body

import "time"

type OrderNumParam struct {
	StartTime time.Time `json:"startTime"`
	EndTime   time.Time `json:"endTime"`
}
