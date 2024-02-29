package main

import (
	"fmt"
	"strconv"
)

func main() {
	sum := 0.0

	loop:
	for {
		var input string

		fmt.Println("Enter a number or type 'proceed' to finish:")
		
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println("Error reading input:", err)
			continue
		}

		if input == "proceed" {
			break
		}

		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Invalid input, enter valid number or type 'proceed'.")
			continue
		}

		sum += num
	}

	fmt.Println("Sum of all numbers is:", sum)
	goto loop
}