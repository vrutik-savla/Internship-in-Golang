package main

import (
	"testing"
	"time"
)

func TestDifferences(t *testing.T) {
	date1 := time.Date(2023, time.December, 20, 11, 23, 56, 00, time.UTC)
	date2 := time.Date(2022, time.November, 15, 10, 14, 30, 00, time.UTC)
	
	gotOutput := difference(date1, date2)
	expectedOutput := []int{1, 35, 1, 9}
	
	for i:=0; i<len(expectedOutput); i++ {
		if expectedOutput[i] != gotOutput[i] {
			t.Errorf("Invalid output: Expected: %d, Got: %d", expectedOutput[i], gotOutput[i])
		}
	}
}
