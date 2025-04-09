package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Phani0310/github-automation-pod/generator" // Update with your module name
)

func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	if err := generator.GenerateDynamicFile(); err != nil {
		log.Fatalf("Error generating file: %v", err)
	}

	if err := runCommand("git", "add", "."); err != nil {
		log.Fatalf("Error adding files: %v", err)
	}

	if err := runCommand("git", "commit", "-m", "Automated file generation"); err != nil {
		log.Fatalf("Error committing changes: %v", err)
	}

	if err := runCommand("git", "push", "origin", "main"); err != nil {
		log.Fatalf("Error pushing to repository: %v", err)
	}

	fmt.Println("File generated and pushed successfully.")
}
