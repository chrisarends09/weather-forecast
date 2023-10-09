package collector

import (
	_ "embed"
	"time"
)

func formatDate(date time.Time, timezone string) string {
	loc, _ := time.LoadLocation(timezone)
	date = date.In(loc)
	return date.Format("01/02/2006")
}

func formatHour(date time.Time, timezone string) string {
	loc, _ := time.LoadLocation(timezone)
	date = date.In(loc)
	// return date.Format("15:04")  this is 24 Hour time
	return date.Format("03:04")
}

func formatTime(date time.Time) string {
	//return date.Format(time.RFC3339)
	return date.Format(time.RubyDate)
}
