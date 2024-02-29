package main

import (
	"testing"
	"time"
)

func TestCountWeekendDays(t *testing.T) {
	parsedDate1, _ := time.Parse("2006-01-02", "2023-12-01")
	parsedDate2, _ := time.Parse("2006-01-02", "2023-12-31")

	gotOutput := countWeekendDays(parsedDate1, parsedDate2)
	expectedOutput := 10
	if gotOutput != expectedOutput {
		t.Errorf("Invalid Output: Expected: %d, Got: %d", expectedOutput, gotOutput)
	}
}