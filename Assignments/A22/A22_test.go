package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCountWorkingDays(t *testing.T) {
	// Define test cases
	testCases := []struct {
		name          string
		currentDate   time.Time
		endDate       time.Time
		holidays      []time.Time
		expectedCount int
	}{
		{
			name:          "NoHolidays",
			currentDate:   time.Date(2023, time.December, 20, 0, 0, 0, 0, time.UTC),
			endDate:       time.Date(2023, time.December, 31, 0, 0, 0, 0, time.UTC),
			holidays:      []time.Time{},
			expectedCount: 8, // Date 23, 24, 30, 31 are weekends
		},
		{
			name:          "WithHolidays",
			currentDate:   time.Date(2023, time.December, 20, 0, 0, 0, 0, time.UTC),
			endDate:       time.Date(2023, time.December, 31, 0, 0, 0, 0, time.UTC),
			holidays:      []time.Time{time.Date(2023, time.December, 25, 0, 0, 0, 0, time.UTC)},
			expectedCount: 7, // Date 23, 24, 30, 31 are weekends and 25 is holiday
		},
	}

	// Iterate through test cases
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Call the function being tested
			actualCount := countWorkingDays(tc.currentDate, tc.endDate, tc.holidays)

			// Check if the actual result matches the expected result
			if !reflect.DeepEqual(actualCount, tc.expectedCount) {
				t.Errorf("Test Case: %s - Expected: %d, Got: %d", tc.name, tc.expectedCount, actualCount)
			}
		})
	}
}
