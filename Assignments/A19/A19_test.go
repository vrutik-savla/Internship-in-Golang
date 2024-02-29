package main

import (
	"testing"
	"time"
)

func TestCompareDate(t *testing.T) {
	parsedDate, _ := time.Parse("2006-01-02", "2023-12-12")

	// Case 1: Check for equal date
	parsedDate1, _ := time.Parse("2006-01-02", "2023-12-12")
	gotResult1 := compareDate(parsedDate, parsedDate1)
	expectedResult1 :=  "Date 1 is equal to Date 2."
	if gotResult1 != expectedResult1 {
		t.Errorf("Invalid Output: Expected: %s, Got: %s", expectedResult1, gotResult1)
	}

	// Case 2: Check for later date
	parsedDate2, _ := time.Parse("2006-01-02", "2023-12-11")
	gotResult2 := compareDate(parsedDate, parsedDate2)
	expectedResult2 := "Date 1 is later than Date 2."
	if gotResult2 != expectedResult2 {
		t.Errorf("Invalid Output: Expected: %s, Got: %s", expectedResult2, gotResult2)
	}

	// Case 3: Check for earlier date
	parsedDate3, _ := time.Parse("2006-01-02", "2023-12-13")
	gotResult3 := compareDate(parsedDate, parsedDate3)
	expectedResult3 := "Date 1 is earlier than Date 2."
	if gotResult2 != expectedResult2 {
		t.Errorf("Invalid Output: Expected: %s, Got: %s", expectedResult3, gotResult3)
	}
}