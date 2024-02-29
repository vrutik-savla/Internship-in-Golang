// Write a program to read the saved HTML file, compress it and store the compressed file as “merce-homepage.html.zip”. Use ready classes for compression. At the end, the program shall print HTML file size, Compressed file size. [file operation concepts]

package main

import (
	"compress/gzip" //To compress file into gzip format
	"fmt"           //For printing
	"io"            //For copying html file content into gzipWriter
	"log/slog"      //For Errors & warnings
	"os"            //For OS operations (Opening/Closing/Reading/Writing)
)

func main() {
	//Retrieve html file size from file statistics
	originalSize, err := getFileStats("merceHomepage.html")
	if err != nil {
		slog.Error("Error getting HTML file stats:", err)
	}
	fmt.Printf("Original HTML file size: %d bytes\n", originalSize)

	// Compress the html file
	filesUnderSize, err := compressAndSaveHTML("merceHomepage.html", "merceHomepage.html.zip")
	if err != nil {
		slog.Error("Error compressing and saving file:", err)
	}
	fmt.Printf("File size under zip: %d bytes\n", filesUnderSize)

	//Retrive file size of "merce-homepage.html.zip".
	compressedSize, err := getFileStats("merceHomepage.html.zip")
	if err != nil {
		slog.Error("Error getting HTML file stats:", err)
	}
	fmt.Printf("Compressed file size: %d bytes\n", compressedSize)

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