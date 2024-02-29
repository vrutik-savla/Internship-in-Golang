// 15. Extend (13) to print time in IST timezone. [Date manipulation, 2 hours]

package main

import (
	"fmt"
	"time"
)

func main() {
	currentTimeIST := time.Now().In(time.FixedZone("IST", 5*60*60+30*60)) //5*60*60+30*60: This calculation represents the offset from UTC in seconds. In this case, it's 5 hours (5*60*60 seconds) and 30 minutes (30*60 seconds) ahead of UTC, totaling 19800 seconds.

	fmt.Println(currentTimeIST.Format("02 Jan 2006"))
	fmt.Println(currentTimeIST.Format("Jan 02, 2006"))
	fmt.Println(currentTimeIST.Format("2006-01-02"))   
	fmt.Println(currentTimeIST.Format("2006-01-02T15:04:05Z"))
	fmt.Println(currentTimeIST.Format("Monday, 02 January 2006"))
}