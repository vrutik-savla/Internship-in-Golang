// Write a program to accept a list of holidays (date in any of the supported formats). Store this list in an internal array. After the user confirms entering of the holiday list, accept a date from the user, and confirm whether it's a working day or not. (All Saturdays and Sundays are implicitly considered as holidays). [Arrays & Date Manipulation, 4 hours].

package main

import (
	"fmt"
	"time"
	"strings"
)

func main() {
	var inputDate string
	var parsedDate time.Time
	var err error

	// User imput for holidays
	holidays := inputHolidays()

	start:
		// User input to check if it's a working day
		fmt.Print("\nEnter a date: ")
		fmt.Scanln(&inputDate)

		// Parse date (string to time for format YYYY-MM-DD)
		parsedDate, err = time.Parse("2006-01-02", inputDate)
		if err != nil {
			fmt.Println("Invalid date format in input:", err)
			goto start
		}

		// Check for working day for date given by user
		if isWorkingDay(parsedDate, holidays) {
			fmt.Println("It's a working day.")
			} else {
			fmt.Println("It's a holiday.")
		}

		inputDate = ""
		goto start
}

// inputHolidays: This function takes no arguments. It initialize empty slice array, takes user input for date in format YYYY-MM-DD & parse it & then stores it in slice array
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

// isWorkingDay: This function takes two arguments, a date in time.Time format & array of time.Time for holidays. It returns false for weekend days & holidays else returns true
func isWorkingDay(date time.Time, holidays []time.Time) bool {
	
	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return false
	}

	for _, holiday := range holidays {
		if holiday.Equal(date) {
			return false
		}
	}

	return true
}