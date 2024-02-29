package main

import (
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Please provide your name as command line argument.")
		return
	}

	name := os.Args[1]
	fmt.Println("Hello", name)
}