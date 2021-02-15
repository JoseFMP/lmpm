package utils

import "time"

type CalendarDay string

type Period struct {
	From time.Time
	To   time.Time
}
