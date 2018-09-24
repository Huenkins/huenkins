package jobstack

import (
	"time"
)

// HistoryPoint is one point for job in history
type HistoryPoint struct {
	ID    string `json:"id"`
	JobID string `json:"job_id"`

	StartTime time.Time `json:"start"`
	Finish    time.Time `json:"finish"`
	Result    string    `json:"result"`
	Server    string    `json:"server"`
}

// Status is current status of one job
type Status struct {
	JobID string `json:"job_id"`

	CreatedTime time.Time      `json:"created_time"`
	Interval    time.Duration  `json:"interval"`
	History     []HistoryPoint `json:"history"`
}
