// Extend (8) to support “sort” operation. Use an in-built function call for sorting numbers. [Core Classes & array operations, 6 hours]

package main

import (
	"fmt"
	"sort"
	"strconv"
)

func main() {
	var numbers[]int

	loop:
	for {
		var input string

		fmt.Println("Enter a number or type sort: ")
		_, err := fmt.Scanln(&input)
		if err != nil {
			fmt.Println("Error reading input.")
			continue
		}

		if input == "sort" {
			break
		}

		num, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Invalid input")
			continue
		}

		numbers = append(numbers, num)
	}

	if len(numbers) > 0 {
		fmt.Println("Unsorted:", numbers)
		sort.Ints(numbers)
		fmt.Println("Sorted:", numbers)
	} else {
	fmt.Println("No numbers to sort.")
	}

	goto loop
}