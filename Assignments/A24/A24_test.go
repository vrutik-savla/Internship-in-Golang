package main

import (
	"testing"
)

// TestCheckNames: This fuction tests checkName that it returns correct boolean or not
func TestCheckNames(t *testing.T) {
	// Create a string map
	strgMap := map[string]string{
		"vrutik": "Vrutik",
		"om": "OM",
	}

	// Case for string exist
	test1 := checkName("vrutik", strgMap)
	if !test1 {
		t.Error("Expected true, got", test1)
	}

	// Case for string does not exist
	test2 := checkName("abc", strgMap)
	if test2 {
		t.Error("Expected false, got", test2)
	}
}