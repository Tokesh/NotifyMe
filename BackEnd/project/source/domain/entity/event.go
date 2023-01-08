package entity

import (
	"time"
)

type Event struct {
	//events_id, event_name, event_timestart, event_timeend, event_result
	Id        string    `json:"events_id"`
	Name      string    `json:"event_name"`
	TimeStart time.Time `json:"event_timestart"`
	TimeEnd   time.Time `json:"event_timeend"`
	Result    string    `json:"event_result"`
}
