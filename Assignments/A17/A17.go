// Write a program to accept two dates (any of the formats supported in the earlier problem) and print a difference in human readable format e.g. “1 year 2 day 32 minutes”. [Date manipulation, 3 hours]

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

		fmt.Println("Enter 1st date (format YYYY-MM-DDT00:00:00Z):")
		fmt.Scanln(&date1)
		fmt.Println("Enter 2nd date (format YYYY-MM-DDT00:00:00Z):")
		fmt.Scanln(&date2)

		parsedDate1, err1 = time.Parse("2006-01-02T15:04:05Z", date1)
		parsedDate2, err2 = time.Parse("2006-01-02T15:04:05Z", date2)
			
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid date format entered.")
			goto start
		}	

		array := difference(parsedDate1, parsedDate2)
		years, days, hours, minutes := array[0], array[1], array[2], array[3]
		fmt.Printf("Difference: %d years %d days %d hours %d minutes\n", years, days, hours, minutes)

		date1, date2 = "", ""
		goto start
}

// difference: This function takes two arguments as date in time.Time format & returns array of difference of years, days, hours & minutes between two dates 
func difference(date1, date2 time.Time) []int {
	difference := date1.Sub(date2)

	years := int(difference.Hours() / 24 / 365)
	days := int(difference.Hours()/24) % 365
	hours := int(difference.Hours()) % 24
	minutes := int(difference.Minutes()) % 60
	
	return []int{years, days, hours, minutes}
}