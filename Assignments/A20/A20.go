// Write a program to accept two dates and print the count of week-end days (Consider Saturdays and Sundays as week-ends). [Loops & Date manipulation & simple expressions, 6 hours]

package main

import (
	"fmt"
	"time"
)

func main() {
	var date1, date2 string
	var parsedDate1, parsedDate2 time.Time
	var err1, err2 error
	
	start:
		// User input for two days
		fmt.Println("Enter 1st date (format YYYY-MM-DD):")
		fmt.Scanln(&date1)
		fmt.Println("Enter 2nd date (format YYYY-MM-DD):")
		fmt.Scanln(&date2)

		// Parse date (Convert string into time with format YYYY-MM-DD)
		parsedDate1, err1 = time.Parse("2006-01-02", date1)
		parsedDate2, err2 = time.Parse("2006-01-02", date2)

		if err1 != nil || err2 != nil {
			fmt.Println("Invalid date format entered.")
			goto start
		}	

		// If 2nd date is before than date 1 then swap dates
		if parsedDate2.Before(parsedDate1) {
			parsedDate1, parsedDate2 = parsedDate2, parsedDate1
		}

		// Count weekend days & print it to user
		weekendCount := countWeekendDays(parsedDate1, parsedDate2)
		fmt.Printf("Number of weekend days between %s & %s: %d\n", date1, date2, weekendCount)

		date1, date2 = "", ""
		goto start
}

// countWeekendDays: This funtion takes two dates in time format, it counts total number of weekend days (Saturday & Sunday) & returns the count value as integer
func countWeekendDays(date1, date2 time.Time) int {
	weekendCount := 0

	for date1.Before(date2) || date1.Equal(date2) {
		if date1.Weekday() == time.Saturday {
			weekendCount += 2
			date1 = date1.AddDate(0, 0, 6)
			continue
		}
		if date1.Weekday() == time.Sunday {
			weekendCount += 1
			date1 = date1.AddDate(0, 0, 5)
			continue
		}
		date1 = date1.AddDate(0, 0, 1)
	}

	return weekendCount
}