// Write a program to accept a filename from command line argument, read a file and print the number of times each word occurs in a file. Perform case insensitive match while counting the occurrence of each word. [Hashmap & File operation & String operation, 6 hours]

package main

import (
	"fmt"      //For printing/scaning
	"log/slog" //For Errors & Warnings
	"os"       //For OS operation (Reading File)
	"strings"  //For String manipulations
)

func main() {
	// Check if file name given or not
	if len(os.Args) < 2 {
		slog.Warn("Please provide a filename as a command line argument.")
		return
	}
	// Store filename
	fileName := os.Args[1]

	//Calling wordOccurrence function which return word count as map
	wordFreq := wordOccurrence(fileName)

	// Print word frequencies
	fmt.Println("Word Frequencies:")
	for word, freq := range wordFreq {
		fmt.Printf("%s: %d\n", word, freq)
	}
}

// wordOccurrence: This function takes a file as argument & return a map where key is each word in the file & value is the number of occurrence of that word
func wordOccurrence(fileName string) map[string]int {
	// Read file, store file contents in (var content []byte) & handle error
	content, err := os.ReadFile(fileName)
	if err != nil {
		slog.Error("Error reading file:", err)
		return nil
	}

	// Convert content of []byte to string, convert all strings to lowercase, split all words by spaces & store it in array slice form into words variable
	words := strings.Fields(strings.ToLower(string(content)))

	// Create map to store word frequencies & count word frequencies
	wordFreq := make(map[string]int)
	for _, word := range words {
		wordFreq[word]++
	}

	return wordFreq
}