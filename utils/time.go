package utils

import "time"

type CalendarDay string

type Period struct {
	From time.Time
	To   time.Time
}

const IanaThailandDatabaseTimeZone = "Asia/Bangkok"

func GetThailandTimeZone() *time.Location {
	location := time.FixedZone(IanaThailandDatabaseTimeZone, 7*60*60) // UTC + 7
	return location
}
