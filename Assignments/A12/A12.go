// Write a program to prompt for two whole positive numbers -- “num1” and “num2”. Print multiplication table for the number e.g. for num1=3 and num2=20, output will be “3 * 1 = 3\n3 * 2 = 6\n … \n3 * 20 = 60”. [For loop, 6 hours]

package main

import (
	"fmt"
)

func main() {
	start:
		var num1, num2 int

		fmt.Println("Enter two numbers:")
		if _, err := fmt.Scanln(&num1, &num2); err != nil {
			fmt.Println("Invalid input:", err)
		}

		for i:=1; i<num2+1; i++ {
			fmt.Println(num1, "*", i, "=", num1*i)
		}
	goto start
}