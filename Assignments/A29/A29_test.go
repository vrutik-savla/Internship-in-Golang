package main

import (
	"os"      //For Reading Writing file
	"reflect" //For comparing two maps
	"testing" //For testing
)

func TestWordOccurrence(t *testing.T) {
	//Create temperory file for testing
	tempFile := "A28Testfile.txt"
	content := []byte("Hello hello world World a cd the when")
	err := os.WriteFile(tempFile, content, 0644)
	if err != nil {
		t.Error(err)
	}
	gotOutput := wordOccurrence(tempFile, []string{"the", "did", "who", "all", "when"})

	//Check if the wordOccurrence returns map as expected.
	expectedOutput := map[string]int{
		"hello":2,
		"world":2,
	}

	// Check if two maps are equal or not & give output
	if !reflect.DeepEqual(gotOutput, expectedOutput) {
		t.Errorf("Unexpected output.\nGot:\n%v\nExpected:\n%v", gotOutput, expectedOutput)
	}

	os.Remove(tempFile)
}