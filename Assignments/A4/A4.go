package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 3 {
		fmt.Println("Please enter your two arguments in command line.")
		return
	}

	template := os.Args[1]
	name := os.Args[2]

	new_template := strings.Replace(template, "<name>", name, -1)

	fmt.Println(new_template)
}