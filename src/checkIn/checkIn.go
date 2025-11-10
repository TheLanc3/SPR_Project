package checkIn

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
	"time"
)

func CheckInDelivery() {
	directoryPath := "../data/"
	destinationDir := "../data/completed"
	entries, err := os.ReadDir(directoryPath)
	if err != nil {
		log.Fatalf("Error reading directory: %v", err)
	}
	// Ensure the destination directory exists
	err = os.MkdirAll(destinationDir, 0755) // Creates the directory and any necessary parent directories
	if err != nil {
		fmt.Printf("Error creating destination directory: %v\n", err)
		return
	}
	// fmt.Printf("Files in directory '%s':\n", directoryPath)
	// Iterate through the directory entries
	for _, entry := range entries {
		// Check if the entry is a file (not a directory)
		if !entry.IsDir() {
			fileName := entry.Name()
			var supplierName string
			charToFind := "."
			index := strings.Index(fileName, charToFind)
			if index != -1 {
				supplierName = fileName[:index]
			} else {
				supplierName = fileName
			}
			now := time.Now()
			formattedTime := now.Format("20060102150405")
			supplierName += charToFind + formattedTime
			sourceFile := filepath.Join(directoryPath, fileName)
			destinationFile := filepath.Join(destinationDir, supplierName)
			err = os.Rename(sourceFile, destinationFile)
			if err != nil {
				fmt.Printf("Error moving file: %v\n", err)
				return
			}
		}
	}
}
