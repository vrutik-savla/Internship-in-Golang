// Write a program to print current date/time in following formats (one line per format) in UTC time.
// [Date manipulation, 3 hours]

package main

import (
	"fmt"
	"time"
)

func main() {
	currentTime := time.Now().UTC()

	// "16 Mar 2022"
	fmt.Println(currentTime.Format("02 Jan 2006"))

	// "Mar 16, 2022"
	fmt.Println(currentTime.Format("Jan 02, 2006"))

	// "2022-03-16"
	fmt.Println(currentTime.Format("2006-01-02"))

	// "2022-03-16T15:52:00Z"
	fmt.Println(currentTime.Format("2006-01-02T15:04:05Z"))

	// "Tuesday, 16 March 2022"
	fmt.Println(currentTime.Format("Monday, 02 January 2006"))

}