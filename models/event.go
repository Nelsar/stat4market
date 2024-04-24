package models

import "time"

type Event struct {
	EventId   int64     `json:"eventId"`
	EventType string    `json:"eventType"`
	UserId    int64     `json:"userID"`
	EventTime time.Time `json:"eventTime"`
	PayLoad   string    `json:"payload"`
}
