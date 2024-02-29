package main

import (
	"os"      //For OS operations
	"testing" //For testing
)

// TestGetFileStats: This function tests the getFileStats function that it returns apropriate size & error or not
func TestGetFileStats(t *testing.T) {
	//Create a temperory test HTML file
	testContent := []byte("Sample content to test the function")
	//Create temperory test file 
	err := createTestFile("test.html", testContent)
	if err != nil {
		t.Error("Error creating test file:", err)
	}

	//Run getFileStats on test HTML file
	size, err := getFileStats("test.html")
	if err != nil {
		t.Error("getFileStats function failed:", err)
	}

	//Check if getFileStats function returns appropriate file size or not
	expectedSize := int64(len(testContent))
	if size != expectedSize {
		t.Errorf("Expected size %d, got %d", expectedSize, size)
	}

	//Testing with a non-existent file, we should get error
	_, err = getFileStats("nonexistent.html")
	if err == nil {
		t.Error("Expected error for non-existent file, but got nil")
	}

	defer os.Remove("test.html")
}

// TestGetFileStats: This function tests the compressAndSaveHTML function that it returns apropriate file size under zip & error or not
func TestCompressAndSaveHTML(t *testing.T) {
	//Create a temperory test HTML file
	content := []byte("Sample content to test the function")
	err := createTestFile("test.html", content)
	if err != nil {
		t.Error("Error creating test HTML file:", err)
	}
	
	//Test case 1: Valid Compression
	sizeUnderZip, err := compressAndSaveHTML("test.html", "test.html.zip")
	testFileSize := int64(len(content))
	if err != nil {
		t.Error("TestCompressAndSaveHTML: Error in valid file compression:", err)
	}
	if sizeUnderZip != testFileSize {
		t.Error("File size before compressing is not equal to file size under zip.")
	}
	defer os.Remove("test.html.zip")
	
	// Test case 2: Non-existent input file
	_, err = compressAndSaveHTML("nonexistent.html", "test.html.zip")
	if err == nil {
		t.Error("TestCompressAndSaveHTML: Expected error for nonexistent input file, but got nil")
	}
	defer os.Remove("test.html.zip")
	
	// Test case 3: Error creating compressed file
	_, err = compressAndSaveHTML("test.html", "/invalid/path/test.html.zip")
	if err == nil {
		t.Error("TestCompressAndSaveHTML: expected error for invalid output path, got nil")
	}

	defer os.Remove("test.html")
}

// createTestFile: This is a helper function for testing purpose which creates a temperory file with provided content in bytes format. It returns error if occured
func createTestFile(fileName string, content []byte) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.Write(content)
	return err
}