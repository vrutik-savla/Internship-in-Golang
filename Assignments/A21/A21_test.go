package main

import (
	"testing"
	"time"
)

func TestIsWorkingDay(t *testing.T) {
	var gotOutput, expectedOutput bool
	// Sample holidays
	holidays := []time.Time{
		time.Date(2023, time.December, 25, 0, 0, 0, 0, time.UTC), // Christmas
	}

	// Test Case 1: For Weekend day
	expectedOutput = false
	gotOutput = isWorkingDay(time.Date(2023, time.December, 23, 0, 0, 0, 0, time.UTC), holidays)
	runTest(gotOutput, expectedOutput, t)
	
	// Test Case 2: For Holiday
	expectedOutput = false
	gotOutput = isWorkingDay(time.Date(2023, time.December, 25, 0, 0, 0, 0, time.UTC), holidays)
	runTest(gotOutput, expectedOutput, t)
	
	// Test Case 3: For Week day
	expectedOutput = true
	gotOutput = isWorkingDay(time.Date(2023, time.December, 27, 0, 0, 0, 0, time.UTC), holidays)	
	runTest(gotOutput, expectedOutput, t)
}

func runTest(gotOutput, expectedOutput bool, t *testing.T) {
	if gotOutput != expectedOutput {
		t.Errorf("Invalid Output: Got: %t, Expected: %t", gotOutput, expectedOutput)
	}
}
