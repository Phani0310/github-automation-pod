package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// GenerateDynamicFile creates a file in /generated with a unique timestamp name
func GenerateDynamicFile() error {
	outputDir := "generated"
	if err := os.MkdirAll(outputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("auto_file_%s.txt", timestamp)
	fullPath := filepath.Join(outputDir, filename)

	content := fmt.Sprintf("This file was generated on %s\n", timestamp)

	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
