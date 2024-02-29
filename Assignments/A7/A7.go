package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	var operation string

	loop:
		for {
			fmt.Println("Enter two numbers and operator: ")
			_, err := fmt.Scanln(&num1, &num2, &operation)
			if err != nil {
				fmt.Println("Error reading input:", err)
				goto loop
			}
			break
		}

	switch operation {
	case "+":
		fmt.Println("num1=", num1, "num2=", num2, "sum=", num1+num2)

	case "-":
		fmt.Println("num1=", num1, "num2=", num2, "difference=", num1-num2)

	case "*":
		fmt.Println("num1=", num1, "num2=", num2, "multiply=", num1*num2)

	case "/":
		if num2 == 0 {
			fmt.Println("Error: Division by zero is not allowed.")
			goto loop
		}
		fmt.Println("num1=", num1, "num2=", num2, "divide=", num1/num2)
		
	default:
		fmt.Println("Invalid Operation Entered")
	}
	goto loop
}