package main

import (
	"testing"
)

// TestSearchString: This function tests searchString that whether it returns valid array or not
func TestSearchString(t *testing.T) {
	var array = []string{"abcd", "abcd efg", "hijklm"}

	gotOutput := searchString("^a", array)
	expectedOutput := []string{"abcd", "abcd efg"}

	if len(gotOutput) != len(expectedOutput) {
        t.Error("Invalid output: Expected:", expectedOutput, "Got:", gotOutput)
    }
	 
	for i := 0; i < len(gotOutput); i++ {
        if gotOutput[i] != expectedOutput[i] {
            t.Error("Invalid output: Expected:", expectedOutput, "Got:", gotOutput)
        }
    }
}