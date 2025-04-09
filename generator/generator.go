// Package generator provides file generation functionality for automation workflows.
package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Generator encapsulates configuration for file generation.
type Generator struct {
	OutputDir string // Directory where the generated files will be stored
}

// NewGenerator returns a new Generator with a default output directory.
func NewGenerator() *Generator {
	return &Generator{
		OutputDir: "generated",
	}
}

// GenerateFile creates a file in the output directory with a unique timestamp-based name.
func (g *Generator) GenerateFile() error {
	// Ensure the output directory exists
	if err := os.MkdirAll(g.OutputDir, os.ModePerm); err != nil {
		return fmt.Errorf("failed to create output directory: %w", err)
	}

	// Create a timestamp-based filename
	timestamp := time.Now().Format("2006-01-02_15-04-05")
	filename := fmt.Sprintf("auto_file_%s.txt", timestamp)

	// Construct the full file path
	fullPath := filepath.Join(g.OutputDir, filename)

	// Create the file content
	content := fmt.Sprintf("This file was generated on %s\n", timestamp)

	// Write the file
	if err := os.WriteFile(fullPath, []byte(content), 0644); err != nil {
		return fmt.Errorf("failed to write file: %w", err)
	}

	return nil
}
