package generator

import (
	"fmt"  // for formatted I/O operations
	"os"   // for file handling ops
	"time" // for working with time stamps
)

// CreateFile generates a file with the specified content
func CreateFile(fileName, content string) error { // file name -> name of filbe to be created and content -> content to be written in file
	// create new file (overwrite if exists)
	file, err := os.Create(fileName)
	if err != nil {
		// returns error if file creation fails
		return err
	}
	defer file.Close() // ensure file is closed when func exists

	// write provided content into file
	_, err = file.WriteString(content)
	if err != nil {
		return err // returns error if writing fails
	}

	fmt.Printf("File created: %s\n", fileName) //
	return nil
}

// generate dynamic file creates a file with a unique name and timestamp
func GenerateDynamicFile() error {
	fileName := fmt.Sprintf("generated_%d.txt", time.Now().Unix())                // generate filename based on current unix stamp
	content := fmt.Sprintf("Generated on: %s\n", time.Now().Format(time.RFC3339)) // create content for file with formatted time stamp
	return CreateFile(fileName, content)                                          // call create file to create file with genrated name and content
}
