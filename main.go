package main

import (
	"fmt"
	"log"     // for logging errors and messages
	"os"      // for accessing system level operations
	"os/exec" // for executing shell commands

	"github.com/Phani0310/github-automation-pod/generator" // importing genrator package
)

// runCommand runs shell commands like git add, commit, and push
func runCommand(command string, args ...string) error {
	cmd := exec.Command(command, args...) // command -> shell command to run (git) and args -> arguments to pass to the command (add)
	cmd.Stdout = os.Stdout                // direct the commands output to the terminal (for debugging)
	cmd.Stderr = os.Stderr                // direct commands error output to the terminal (for debugging)
	return cmd.Run()
}

func main() {
	// generate a file using the generator package (create a file with a unique timestamp based name and content)
	if err := generator.GenerateDynamicFile(); err != nil {
		log.Fatalf("Error generating file: %v", err) // if file genration fails, log the error and exit
	}

	// stage the changes using git add . stages all modified files for commit
	if err := runCommand("git", "add", "."); err != nil {
		log.Fatalf("Error adding files: %v", err)
	}

	// commit the changes with custom message and || operator will prevent errors if no changes are made
	if err := runCommand("git", "commit", "-m", "Automated file generation"); err != nil {
		log.Fatalf("Error committing changes: %v", err)
	}

	// push changes to the remote repository and pushes to main "branch of origin" repository
	if err := runCommand("git", "push", "origin", "main"); err != nil {
		log.Fatalf("Error pushing to repository: %v", err)
	}

	fmt.Println("File generated and pushed successfully.") // if everything successful, prints given message
}
