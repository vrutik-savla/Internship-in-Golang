//Same as (21) to accept a list of holidays, and then prompt a user for two dates (in the supported format) and print the number of working days between two dates. Consider both dates during the calculation. [Date manipulation & loop & boolean expressions, 4 hours]

package main

import (
	"fmt"
	"strings"
	"time"
)

func main() {
	var startDate, endDate string
	var parsedStartDate, parsedEndDate time.Time
	var err1, err2 error

	// User input for holidays
	holidays := inputHolidays()

	start:
	// User input of two dates to find working days between them
	fmt.Print("Enter the first date (YYYY-MM-DD format): ")
	fmt.Scanln(&startDate)
	fmt.Print("Enter the second date (YYYY-MM-DD format): ")
	fmt.Scanln(&endDate)

	// Parsing of two dates
	parsedStartDate, err1 = time.Parse("2006-01-02", startDate)
	parsedEndDate, err2 = time.Parse("2006-01-02", endDate)
	if err1 != nil || err2 != nil {
		fmt.Println("Invalid Input:", err1, err2)
		goto start
	}

	// Count number of working days between given two dates by user
	workingDays := countWorkingDays(parsedStartDate, parsedEndDate, holidays)
	fmt.Printf("\nNumber of working days between %s and %s: %d\n", startDate, endDate, workingDays)

	startDate, endDate = "", ""
	goto start
}

// inputHolidays: This function takes no arguments. It initialize empty slice array, takes user input for date in format YYYY-MM-DD & parse it & then stores it in slice array.
func inputHolidays() []time.Time {
	var holidays []time.Time
	
	fmt.Print("Enter list of holidays (Format YYYY-MM-DD). Press Enter after each date. Type 'Done' to finish:\n")
	
	for {
		var holiday string
		fmt.Scanln(&holiday)

		if strings.ToLower(holiday) == "done" {
			break
		}
		parsedTime, err := time.Parse("2006-01-02", holiday)
		if err == nil {
			holidays = append(holidays, parsedTime)	
		} else {
			fmt.Println("Entered date does not match YYYY-MM-DD format.")
			continue
		}
	}
	return holidays
}

// countWorkingDays: This function takes two dates & array of holidays all in tine.Time format. It counts total number of working days between given two dates & returns the count as integer value.
func countWorkingDays(currentDate, endDate time.Time, holidays []time.Time) int {
	workingDays := 0

	for currentDate.Before(endDate) || currentDate.Equal(endDate) {
		if currentDate.Weekday() == time.Friday {
			currentDate = currentDate.AddDate(0, 0, 2)
			continue
		}

		for _, holiday := range holidays {
			if holiday.Equal(currentDate) {
				currentDate = currentDate.AddDate(0, 0, 1)
				continue
			}
		}
		workingDays += 1
		currentDate = currentDate.AddDate(0, 0, 1)
	}

	return workingDays
}