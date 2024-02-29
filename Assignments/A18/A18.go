// Write a program to accept a date and print whether the date falls into a leap year. Accept a date in any format supported in one of the previous problems. [Date manipulation, 2 hours]

package main

import (
	"fmt"
	"time"
)

func main() {
	var inputDate string
	var parsedDate time.Time
	var err error
	start:
		// User input
		fmt.Println("Enter a date (format YYYY-MM-DD): ")
		fmt.Scanln(&inputDate)

		//Take input string & parse date in layout YYYY-MM-DD & store it as time data type
		parsedDate, err = time.Parse("2006-01-02", inputDate)
		if err != nil {
			fmt.Println("Invalid date format entered.")
			goto start
		}	

		//Get year from given date by user as integer
		year := parsedDate.Year()

		// Check for leap year
		if isLeapYear(year) {
			fmt.Println(year, "is a leap Year")
			} else {
			fmt.Println(year, "is not a leap Year")
		}
		goto start
}

func isLeapYear(year int) bool {
	return (year%4 == 0 && year%100 != 0) || (year%400 ==0 )
}