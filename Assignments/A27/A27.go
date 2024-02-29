// Extend (26) to accept URL as a command line argument instead of a hardcoded URL within the program. [Revision of previous concepts, 2 hours]
// https://natours-vrutik.netlify.app

package main

import (
	"compress/gzip" //To compress file into gzip format
	"fmt"           //For printing
	"io"            //For copying html file content into gzipWriter
	"log/slog"
	"strings" //For string manipulations

	//For Errors & warnings
	"net/http"
	"os" //For OS operations (Opening/Closing/Reading/Writing)
)

func main() {
	//Get URL from command line argument, if not present throw an error
	if len(os.Args) < 2 {
		fmt.Println("Please provide a URL as command line argument.")
		return
	}
	url := os.Args[1]

	//Save HTML content of provided URL by creating HTML file
	fileName, err:= saveHTMLFromURL(url)
	if err != nil {
		fmt.Println("Error occured in getting respose, Please provide valid URL:", err)
		return
	}

	//Retrieve html file size from file statistics
	originalSize, err := getFileStats(fileName)
	if err != nil {
		slog.Error("Error getting HTML file stats:", err)
	}
	fmt.Printf("Original HTML file size: %d bytes\n", originalSize)

	// Compress the html file
	zipFileName := fileName + ".zip"
	filesUnderSize, err := compressAndSaveHTML(fileName, zipFileName)
	if err != nil {
		slog.Error("Error compressing and saving file:", err)
	}
	fmt.Printf("File size under zip: %d bytes\n", filesUnderSize)

	//Retrive file size of "merce-homepage.html.zip".
	compressedSize, err := getFileStats(zipFileName)
	if err != nil {
		slog.Error("Error getting HTML file stats:", err)
	}
	fmt.Printf("Compressed file size: %d bytes\n", compressedSize)
}

//saveHTMLFromURL: This function takes a URL, creates HTML file & saves URL's HTML content to the file Returns created HTML filename & error if occurred
func saveHTMLFromURL(url string) (string, error) {
	//Make HTTP GET request to URL 
	response, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer response.Body.Close()

	//Extract filename from URL
	fileName := extractFileNameFromURL(url) +".html"

	//Create new file to save HTML Content
	file, err := os.Create(fileName)
	if err != nil {
		return "", err
	}
	defer file.Close()

	// Copy the HTML content from the response body to the file
	_, err = io.Copy(file, response.Body)
	if err != nil {
		return "", err
	}

	fmt.Println("HTML content saved to", fileName)
	return fileName, nil
}

// extractFileNameFromURL: This function takes a url & creates a valid filename w.r.t that URL
func extractFileNameFromURL(url string) string {
	//Split URL by '/'
	parts := strings.Split(url, "/")

	//Get last part of url
	lastPart := parts[len(parts) - 1]

	//Check if last part contains any queries
	if strings.Contains(lastPart, "?") {
		//Split by '?' & remove queries
		fileNameParts := strings.Split(lastPart, "?") 
		return fileNameParts[0]
	}

	return lastPart
}

// getFileStats: This function returns file size under zip and error as per given file as argument
func getFileStats(fileName string) (int64, error) {
	//Open the file, file is pointer pointing to this file
	file, err := os.Open(fileName)
	if err != nil {
		return 0, err
	}
	defer file.Close()

	//Store file statistics in fileStat
	fileStat, err := file.Stat()
	if err != nil {
		return 0, err
	}

	return fileStat.Size(), nil
}

// compressAndSaveHTML: This function takes a two arguments(file to be compressed, compressed file name). It creates a new zip file with given name & return file size under zip or error if occurred
func compressAndSaveHTML(inputFileName, outputFileName string) (int64, error) {
	//Open the file, file is pointer pointing to this file
	file, err := os.Open(inputFileName)
	if err != nil {
		return -1, err
	}
	defer file.Close()

	// Create a new empty file, gzipFile is pointer pointing to this new File
	gzipFile, err := os.Create(outputFileName)
	if err != nil {
		return -1, err
	}
	defer gzipFile.Close()

	//Initialize gzip writer to compress data & write content to created file
	gzipWriter := gzip.NewWriter(gzipFile)
	defer gzipWriter.Close()

	//Copies the HTML file's content to the gzip.Writer (which compresses it) and writes the compressed data to "merce-homepage.html.zip".
	compressedSize, err := io.Copy(gzipWriter, file)
	if err != nil {
		return -1, err
	}

	return compressedSize, nil
}
