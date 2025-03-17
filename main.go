package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Phani0310/github-automation-pod/generator"
)

// runCommand runs shell commands like git add, commit, and push
func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	// Step 1: Generate a file using the generator package
	if err := generator.GenerateDynamicFile(); err != nil {
		log.Fatalf("Error generating file: %v", err)
	}

	// Step 2: Stage the changes
	if err := runCommand("git", "add", "."); err != nil {
		log.Fatalf("Error adding files: %v", err)
	}

	// Step 3: Commit the changes
	if err := runCommand("git", "commit", "-m", "Automated file generation"); err != nil {
		log.Fatalf("Error committing changes: %v", err)
	}

	// Step 4: Push changes to the remote repository
	if err := runCommand("git", "push", "origin", "main"); err != nil {
		log.Fatalf("Error pushing to repository: %v", err)
	}

	fmt.Println("✅ File generated and pushed successfully.")
}
