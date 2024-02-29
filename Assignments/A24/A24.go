// Write a program to take the names of candidates as input. Prompt user to keep entering new names till user enters “done”. Once a list of names are accepted, prompt the user for a name. Output shall be “<name> exists” or “<name> does not exist”. A name exists if the name exactly matches one of the names provided earlier. Use case insensitive match for comparison. [Hash data structure & String operations, 6 hours]

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	// Initialise Scanner
	scanner := bufio.NewScanner(os.Stdin)

	// Initialise empty map
	names := make(map[string]string)
	for {
		fmt.Println("Enter a name or type 'Done' to finish. (Press Enter after each name):")
		scanner.Scan()
		// Store names into name variable
		name := scanner.Text()
		if strings.ToLower(name) == "done" {
			// When done, break the loop
			break
		}
		// Store name into map
		names[strings.ToLower(name)] = name
	}
	
	var name string
	loop:
	// User input to check name
	fmt.Println("Enter a name to check if it exists (Press Enter to exit): ")
	scanner.Scan()
	name = scanner.Text()

	if name == "" {
		return
	} else if checkName(name, names) {
		fmt.Printf("%s exists.\n", name)
		goto loop
	} else {
		fmt.Printf("%s does not exist.\n", name)
		goto loop
	}
}

// checkName: This function takes string & a map of strings. It returns true if string exist in map else false. (Case insensitive)
func checkName(strg string, strgMap map[string]string) bool{
	if _, exists := strgMap[strings.ToLower(strg)]; exists {
		return true
	} else {
		return false
	}
}