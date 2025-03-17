package generator

import (
	"fmt"
	"os"
	"time"
)

// CreateFile generates a file with the specified content
func CreateFile(fileName, content string) error {
	file, err := os.Create(fileName)
	if err != nil {
		return err
	}
	defer file.Close()

	_, err = file.WriteString(content)
	if err != nil {
		return err
	}

	fmt.Printf("File created: %s\n", fileName)
	return nil
}

// GenerateDynamicFile creates a file with a unique name and timestamp
func GenerateDynamicFile() error {
	fileName := fmt.Sprintf("generated_%d.txt", time.Now().Unix())
	content := fmt.Sprintf("Generated on: %s\n", time.Now().Format(time.RFC3339))
	return CreateFile(fileName, content)
}
