package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanf("%v", &name)
	fmt.Println("Hello", name)
}