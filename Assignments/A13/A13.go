// Write a program to prompt for three inputs: character to be used for display, num1 to represent number of rows and num2 to represent number of columns. The output will be a rectangular matrix where each cell will print a character input as a first input value. [Loops, 4 hours]

package main

import (
	"fmt"
)

func main() {
	start:
		var character string
		var row, column int

		fmt.Println("Enter character, row & column:")
		if _, err := fmt.Scanln(&character,&row,&column); err != nil {
			fmt.Println("Invalid Input:", err)
		}

		for i:=0; i<row; i++ {
			for j:=0; j<column; j++ {
				fmt.Printf("%s\t", character)
			}
			fmt.Println()
		}

	goto start
}
