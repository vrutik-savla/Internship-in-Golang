// Extend the above program to ignore common words (e.g. “the”, “a”, “an”, …) and single letter words (e.g. “I”). [Revision of previous concepts, 3 hours]

package main

import (
	"fmt"      //For printing
	"log/slog" //For printing warning & error
	"os"       //For OS operation (Reading / Writing)
	"strings"  //For string manipulations
)

func main() {
	// Checking if filename given or not
	if len(os.Args) < 2 {
		slog.Warn("Please provide a filename as a command-line argument.")
		return
	}

	// Storing the file name
	fileName := os.Args[1]	

	//Calling wordOccurrence function which return word count as map
	wordFreq := wordOccurrence(fileName, ignoreStrings())

	// Print word frequencies
	fmt.Println("Word frequencies (excluding common and two-letter words):")
	for word, freq := range wordFreq {
		fmt.Printf("%s: %d\n", word, freq)
	}
}

// wordOccurrence: This function takes a file name & array of strings which will be ignored, as argument & returns a map where key is each word in the file & value is the number of occurrence of that word. (It ignores words with atmost two alphabets & words which are given in array as 2nd argument)
func wordOccurrence(fileName string, ignoreArray []string) map[string]int {
	//Read file, store file contents in (var content []byte) & handle error
	content, err := os.ReadFile(fileName)
	if err != nil {
		slog.Error("Error opening file:", err)
		return nil
	}

	// Convert content of []byte to string, convert all strings to lowercase, split all words by spaces & store it in array slice form into words variable
	words := strings.Fields(strings.ToLower(string(content)))

	//Converting ignore list from array to map
	ignoreList := make(map[string]bool)
	for _, value := range ignoreArray {
		ignoreList[value] = true
	}

	// Create map to store word frequencies & count word frequencies
	wordFreq := make(map[string]int)
	for _, word := range words {
		if len(word) > 2 && !ignoreList[word] {
			wordFreq[word]++
		}
	}

	return wordFreq
}

// Ignore list for wordOccurrence function
func ignoreStrings() []string {
	return []string{"the", "did", "who", "all", "when"}
}