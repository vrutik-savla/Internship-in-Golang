// Write a program to accept two dates (any of the supported period) and print an output whether date1 and date2 are equal, date1 is earlier than date2 or date1 is later than date2. [Date comparison, 1 hours]

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
		// User input for two dates
		fmt.Println("Enter two dates with space in between (format YYYY-MM-DD): ")
		fmt.Scanln(&date1, &date2)

		// Parse two dates (Convert string to time & check for valid layout)
		parsedDate1, err1 = time.Parse("2006-01-02", date1)
		parsedDate2, err2 = time.Parse("2006-01-02", date2)

		fmt.Println(parsedDate1, parsedDate2)

		// If error occurs, send apropriate message 
		if err1 != nil || err2 != nil {
			fmt.Println("Invalid Input:", err1, err2)
			goto start
		}	

		getComparedStrg := compareDate(parsedDate1, parsedDate2)
		fmt.Println(getComparedStrg)
		
		date1, date2 = "", ""
		goto start
}

// compareDate: This fuction takes two arguments dates in time.Time format & compares two dates & returns a string w.r.t comparison
func compareDate(date1, date2 time.Time) string {
	switch {
		case date1.Before(date2):
			return "Date 1 is earlier than Date 2."
		case date1.After(date2):
			return "Date 1 is later than Date 2."
		case date1.Equal(date2):
			return "Date 1 is equal to Date 2."
		default:
			return "Invalid"
	}
}