// Accept a list of holidays, and then prompt a user for two inputs: input date as a first argument and a number of business days as a second argument. Number of business days will be a positive or negative whole number. The output shall be the date relative to an input date +/- the number of business days. Holidays must be excluded while calculating the output date.

package main

import (
	"fmt"
	"time"
)

func main() {
	var inputDate string
	var parsedInputDate time.Time
	var err error
	var businessDays int

	// Take user input for holidays
	holidays := inputHolidays()

	start:
	fmt.Println("Enter date: ")
	fmt.Scanln(&inputDate)
	parsedInputDate, err = time.Parse("2006-01-02", inputDate)
	if err != nil {
		fmt.Println("Invalid Input:", err)
		goto start
	}

	fmt.Println("Enter number of business days(+ / -):")
	fmt.Scanln(&businessDays)

	outputDate := calculateBusinessDay(parsedInputDate, businessDays, holidays)
	fmt.Printf("The new date after %d business days from %s is: %s\n", businessDays, parsedInputDate.Format("2006-01-02"), outputDate.Format("2006-01-02"))

	inputDate = ""
	goto start
}

// inputHolidays: This function takes no arguments. It initialize empty slice array, takes user input for date in format YYYY-MM-DD & parse it & then stores it in slice array.
func inputHolidays() []string {
	var holidays []string

	fmt.Println("Enter list of holidays (format: YYYY-MM-DD). Press 'Enter' after each date. Type 'Done' to finish.")

	for {
		var holiday string
		fmt.Scanln(&holiday)
		if holiday == "done" {
			break
		}
		holidays = append(holidays, holiday)
	}
	return holidays
}

// calculateBusinessDay: This function takes a date (time.Time), number of business days (integer) & list of holdidays (array of strings). It returns new date (time.Time) after the count of that business days. 
func calculateBusinessDay(startDate time.Time, businessDays int, holidays []string) time.Time {
	currentDate := startDate
	remainingBusinessDays := businessDays

	for remainingBusinessDays != 0 {
		currentDate = currentDate.AddDate(0, 0, businessDays/abs(businessDays))

		if isWorkingDay(currentDate, holidays) {
			remainingBusinessDays -= businessDays / abs(businessDays)
		}
	}
	return currentDate
}

// isWorkingDay: This function takes a date as time.Time & list of holidays as array of strings & returns a boolean. If date is Saturday, SUnday or matched with date of holiday then it returns false, else returns true.
func isWorkingDay(date time.Time, holidays []string) bool {
	if date.Weekday() == time.Saturday || date.Weekday() == time.Sunday {
		return false
	}
	for _, holiday := range holidays {
		holidayDate, err := time.Parse("2006-01-02", holiday)
		if err == nil && holidayDate.Equal(date) {
			return false
		}
	}
	return true
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}