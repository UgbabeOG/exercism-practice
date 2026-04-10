package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
// Layout: "1/2/2006 15:04:05"
func Schedule(date string) time.Time {
	t, _ := time.Parse("1/2/2006 15:04:05", date)
	return t
}

// HasPassed returns whether a date has passed.
// Layout: "January 2, 2006 15:04:05"
func HasPassed(date string) bool {
	t, _ := time.Parse("January 2, 2006 15:04:05", date)
	return t.Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
// Layout: "Monday, January 2, 2006 15:04:05"
// Afternoon is defined as 12:00 to 18:00 (exclusive of 18:00).
func IsAfternoonAppointment(date string) bool {
	t, _ := time.Parse("Monday, January 2, 2006 15:04:05", date)
	hour := t.Hour()
	return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
// Output format: "You have an appointment on Thursday, July 25, 2019, at 13:45."
func Description(date string) string {
	t := Schedule(date)
	return fmt.Sprintf("You have an appointment on %s.", t.Format("Monday, January 2, 2006, at 15:04"))
}

// AnniversaryDate returns a Time with this year's anniversary (Sept 15th).
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC)
}
