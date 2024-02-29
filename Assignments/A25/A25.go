// Write a program to take the names of candidates as input. Prompt user to keep entering new names till user enters “done”. Once a list of names are accepted, prompt the user for a search pattern (regex). Output shall be a list of all names where the search pattern exists. [Array & Loop & Regex, 8-12 hours]

package main

import (
	"bufio" //For scanning
	"fmt" //For printing
	"os" //For OS operations
	"regexp" //For regular expression search
	"strings" //For string operations
)

func main() {
	var names []string
	
	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Println("Enter Name or type 'Done' to finish: ")
		scanner.Scan()
		name := scanner.Text()
		if strings.ToLower(name) == "done" {
			break
		}

		names = append(names, name)
	}

	loop:
	fmt.Println("Enter search pattern (regular expression): ")
	scanner.Scan()
	searchPattern := scanner.Text()
	matchedArray := searchString("(?i)"+searchPattern, names)

	if len(matchedArray) == 0 {
		fmt.Println("No Match Found")
	} else {
		for i := 0; i < len(matchedArray); i++ {
			fmt.Println(matchedArray[i])
		}
	}
	searchPattern = ""
	goto loop
}

// searchString: This function takes a regular expression as string & array of strings. It prints all the strings present in array that satisfies the given regular expression
func searchString(regExp string, arrayString []string) []string {
	var matchedArray []string
	regExpP, err := regexp.Compile(regExp)
	if err != nil {
		fmt.Println("Invalid Regular Expression:", err)
		return []string{"Try Again"}
	} else {
		for _, name := range arrayString {
			if regExpP.MatchString(name) {
				matchedArray = append(matchedArray, name)
			}
		}
	}
	return matchedArray
}

/*
Pattern: john
Description: Matches the exact string "john".
Case-insensitive search:

Pattern: (?i)mary
Description: Performs a case-insensitive search for "mary".
Matching names starting with a specific letter:

Pattern: ^a
Description: Matches names starting with the letter 'a'.
Matching names ending with a specific letter:

Pattern: z$
Description: Matches names ending with the letter 'z'.
Matching names containing a specific substring:

Pattern: .*anna.*
Description: Matches names containing "anna" anywhere within the name.
Matching names of a specific length:

Pattern: ^.{4}$
Description: Matches names that are exactly 4 characters long.
Matching names with specific characters:

Pattern: [aeiou]
Description: Matches names containing any vowel ('a', 'e', 'i', 'o', 'u').
*/

/*
x*             zero or more x, prefer more
x+             one or more x, prefer more
x?             zero or one x, prefer one
x{n,m}         n or n+1 or ... or m x, prefer more
x{n,}          n or more x, prefer more
x{n}           exactly n x
x*?            zero or more x, prefer fewer
x+?            one or more x, prefer fewer
x??            zero or one x, prefer zero
x{n,m}?        n or n+1 or ... or m x, prefer fewer
x{n,}?         n or more x, prefer fewer
x{n}?          exactly n x
*/