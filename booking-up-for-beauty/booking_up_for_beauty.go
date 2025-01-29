package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	layouts := []string{
		"January 2, 2006 15:04:05", 
		"1/2/2006 15:04:05",
		"Monday, January 2, 2006 15:04:05",
	}

	for _, layout := range layouts {
		t, err := time.Parse(layout, date)
		if err == nil {
			return t
		}
	}
	fmt.Println("Error: Unable to parse date with any known format")
	return time.Time{}
}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	return Schedule(date).Before(time.Now())
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	return Schedule(date).Hour() >= 12 && Schedule(date).Hour() < 18 
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	return fmt.Sprintf("You have an appointment on %v", Schedule(date).Format("Monday, January 2, 2006, at 15:04."))
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	return time.Date(time.Now().Year(), time.September, 15, 0, 0, 0, 0, time.UTC )
}
