package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numbers []float64
	loop:
		for {
			var input string

			fmt.Println("Enter a number or operation(count, mean, minimum, maximum):")
			if _, err := fmt.Scanln(&input); err != nil {
				fmt.Println("Error reading input:", err)
				continue
			}

			switch input {
			case "count":
				count := count(numbers)
				fmt.Println("Count of numbers entered:", count)

			case "mean":
				mean := mean(numbers)
				fmt.Println("Mean of numbers entered:", mean)

			case "minimum":
				minimum := minimum(numbers)
				fmt.Println("Minimum of entered number:", minimum)
			
			case "maximum":
				maximum := maximum(numbers)
				fmt.Println("maximum of entered number:", maximum)

			default:
				num, err := strconv.ParseFloat(input, 64)
				if err != nil {
					fmt.Println("Invalid input.")
					continue
				}
				numbers = append(numbers, num)
			}
			goto loop
		}
}

func count(nums []float64) int {
	return len(nums)
}

func mean(nums []float64) float64 {
	sum := 0.0
	for _, num := range nums {
		sum += num
	}
	if len(nums) == 0 {
		return 0
	}
	return sum / float64(len(nums))
}

func minimum(nums []float64) float64 {
	min := nums[0]
	for _, num := range nums {
		if num < min {
			min = num
		}
	}
	return min
}

func maximum(nums []float64) float64 {
	max := nums[0]
	for _, num := range nums {
		if num > max {
			max = num
		}
	}
	return max
}