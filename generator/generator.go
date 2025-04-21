package generator

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
)

// Generator handles dynamic file creation logic
type Generator struct {
	OutputDir string
}

// New creates a new Generator with a default output dir
func New() *Generator {
	return &Generator{OutputDir: "generated"}
}

// Generate creates a new file with current timestamp content
func (g *Generator) Generate() (string, error) {
	if err := os.MkdirAll(g.OutputDir, os.ModePerm); err != nil {
		return "", fmt.Errorf("unable to create output directory: %w", err)
	}

	ts := time.Now().Format("2006-01-02_15-04-05")
	file := filepath.Join(g.OutputDir, fmt.Sprintf("gen_%s.txt", ts))
	content := fmt.Sprintf("Generated at: %s\n", ts)

	if err := os.WriteFile(file, []byte(content), 0644); err != nil {
		return "", fmt.Errorf("unable to write file: %w", err)
	}

	return file, nil
}
