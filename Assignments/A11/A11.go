// Extend (9) to support “countodd” and “counteven” operations to respectively print number of times odd numbers and number of even numbers found within the list.
package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numbers []int
	
	loop:
	even := 0
	odd := 0
	for {
			var input string

			fmt.Println("Enter an integer or proceed: ")
			if _,err := fmt.Scanln(&input); err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			if input == "proceed" {
				
				for _, num := range numbers {
					if num%2 == 0 {
						even += 1
					} else {
						odd += 1
					}
				}
				
				fmt.Println("Total odd numbers are:", odd, "and even numbers are:", even)
			} else {
				
				num, err := strconv.Atoi(input)
				if err != nil {
					fmt.Println("Invalid input")
					continue
				}
				numbers = append(numbers, num)
			
			}
			goto loop
		}
}