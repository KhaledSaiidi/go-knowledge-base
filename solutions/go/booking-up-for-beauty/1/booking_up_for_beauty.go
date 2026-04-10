package booking
// Time 	Options
// Year 	2006 ; 06
// Month 	Jan ; January ; 01 ; 1
// Day 	02 ; 2 ; _2 (For preceding 0)
// Weekday 	Mon ; Monday
// Hour 	15 ( 24 hour time format ) ; 3 ; 03 (AM or PM)
// Minute 	04 ; 4
// Second 	05 ; 5
// AM/PM Mark 	PM
// Day of Year 	002 ; __2
// Layout: Jan 2 15:04:05 2006 -> 1/2/2006 15:04:05
import "time"
import "fmt"

// Schedule returns a time.Time from a string containing a date.
func Schedule(date string) time.Time {
    layout := "1/2/2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        panic(err) // or handle properly
    }
    return t}

// HasPassed returns whether a date has passed.
func HasPassed(date string) bool {
    layout := "January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        panic(err) // or handle properly
    }
	return time.Now().After(t)
}

// IsAfternoonAppointment returns whether a time is in the afternoon.
func IsAfternoonAppointment(date string) bool {
    layout := "Monday, January 2, 2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        panic(err)
    }
    hour := t.Hour()
    return hour >= 12 && hour < 18
}

// Description returns a formatted string of the appointment time.
func Description(date string) string {
    layout := "1/2/2006 15:04:05"
    t, err := time.Parse(layout, date)
    if err != nil {
        panic(err)
    }
	s := fmt.Sprintf("You have an appointment on %s, %s %d, %d, at %02d:%02d.", t.Weekday(), t.Month(), t.Day(), t.Year(), t.Hour(), t.Minute())
	return s
}

// AnniversaryDate returns a Time with this year's anniversary.
func AnniversaryDate() time.Time {
    t := time.Date(2026,time.September,15,00,0,0,0,time.UTC)
    return t
}

