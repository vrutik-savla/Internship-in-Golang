package main

import (
	"reflect"
	"testing"
	"time"
)

func TestCalculateBusinessDay(t *testing.T) {
	startDate := time.Date(2023, time.December, 20, 0, 0, 0, 0, time.UTC)
	holidays := []string{"2023-12-25", "2023-12-31"}

	testCases := []struct {
		name string
		startDate time.Time
		businessDays int
		holidays []string
		expectedResult time.Time
	}{
		{
			name: "PositiveBusinessDays_NoHolidays",
			startDate: startDate,
			businessDays: 5,
			holidays: []string{},
			expectedResult: time.Date(2023, time.December, 27, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "PositiveBusinessDays_WithHolidays",
			startDate: startDate,
			businessDays: 5,
			holidays: holidays,
			expectedResult: time.Date(2023, time.December, 28, 0, 0, 0, 0, time.UTC),
		},
		{
			name: "NegativeBusinessDays_NoHolidays",
			startDate: startDate,
			businessDays: -5,
			holidays: []string{},
			expectedResult: time.Date(2023, time.December, 13, 0, 0, 0, 0, time.UTC),
		},
		{
			name:           "NegativeBusinessDays_WithHolidays",
			startDate:      startDate,
			businessDays:   -5,
			holidays:       holidays,
			expectedResult: time.Date(2023, time.December, 13, 0, 0, 0, 0, time.UTC),
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			gotResult := calculateBusinessDay(tc.startDate, tc.businessDays, tc.holidays)
			if !reflect.DeepEqual(gotResult, tc.expectedResult) {
				t.Errorf("Test case: %s, Expected: %s, Got: %s", tc.name, tc.expectedResult, gotResult)
			}
		})
	}
}

func TestIsWorkingDay(t *testing.T) {
	testDate := time.Date(2023, time.December, 25, 0, 0, 0, 0, time.UTC)
	holidays := []string{"2023-12-25", "2023-12-31"}

	testCases := []struct {
		name string
		testDate time.Time
		holidays []string
		expectedResult bool
	}{
		{
			name: "Holiday",
			testDate: testDate,
			holidays: holidays,
			expectedResult: false,
		},
		{
			name: "RegularWorkingDay",
			testDate: time.Date(2023, time.December, 20, 0, 0, 0, 0, time.UTC),
			holidays: holidays,
			expectedResult: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actualResult := isWorkingDay(tc.testDate, tc.holidays)
			if !reflect.DeepEqual(actualResult, tc.expectedResult) {
				t.Errorf("Test Case: %s, Expected: %t, Got: %t", tc.name, tc.expectedResult, actualResult)
			}
		})
	}
}