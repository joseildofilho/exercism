package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	value, _ := time.Parse("1/02/2006 15:04:05", date)
	return value
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	value, err := time.Parse("January 2, 2006 15:05:05", date)
	if err != nil {
		return false
	}
	return value.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	value, err := time.Parse("Monday, January 2, 2006 15:04:05", date)
	if err != nil {
		return false
	}
	return value.Hour() >= 12  && value.Hour() < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	value, _ := time.Parse("1/2/2006 15:04:05", date)
	return fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %d:%d.", value.Weekday(), value.Month(), value.Day(), value.Year(), value.Hour(), value.Minute())
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
