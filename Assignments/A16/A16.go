// Print supported timezones, and accept a valid timezone as input and print time as per the time zone selected by an end user. [Date manipulation & switch statement, 3 hours]

package main

import (
	"fmt"
	"time"

	"example.com/timezones"
)

func main() {
	//Import supported timezones from external package
	supportedTimeZones := timezones.Timezones()
	
	// Print supported timezones
	fmt.Println("Supported Timezones:")
	for _, tz := range supportedTimeZones {
		fmt.Println(tz)
	}
	fmt.Println()
	

	input:
	var userInput string
	// User input
	fmt.Println("Enter a timezone: ")
	fmt.Scanln(&userInput)
	
	// Get location of timezone provided by user
	loc, err := time.LoadLocation(userInput)
	if err != nil {
		fmt.Println("Error loading location, Please provide valid timezone:", err)
		goto input
	}

	// Load time of that location
	timeInZone := time.Now().In(loc)
	fmt.Printf("Time in %s: %s\n", userInput, timeInZone.Format("2006-01-02 15:04:05"))
	goto input
}
