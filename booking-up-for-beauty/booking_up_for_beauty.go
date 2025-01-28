package booking

import (
	"fmt"
	"time"
)

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
	
	layout := "1/2/2006 15:04:05"
	
	t, err := time.Parse(layout, date)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Time{}
	}

	return t
}

func hasPassedTime(date string) time.Time {
	layout := "January 2, 2006 15:04:05"
	t, err := time.Parse(layout, date)

	if err != nil {
		fmt.Println("Error parsing date:", err)
		return time.Time{}
	}
	return t

}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
	t := hasPassedTime(date).Before(time.Now())
	return t
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
	panic("Please implement the IsAfternoonAppointment function")
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
	panic("Please implement the Description function")
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
	panic("Please implement the AnniversaryDate function")
}
