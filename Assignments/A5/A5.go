package main

import "fmt"

func main() {
	var num1, num2 int

	Loop:
		for {
			fmt.Println("Enter two integers:")
			if _, err := fmt.Scanln(&num1, &num2); err != nil {
				fmt.Println("Error:", err)
				goto Loop
			}
			break
		}

	fmt.Println("num1=", num1, "num2=", num2, "sum=", num1+num2)
}