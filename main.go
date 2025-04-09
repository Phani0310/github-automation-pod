// Command main runs the GitHub automation pod to generate and push files automatically.
package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"

	"github.com/Phani0310/github-automation-pod/generator" // Replace with your actual module path
)

// GitRunner provides an interface for executing Git operations.
type GitRunner struct{}

// run executes a shell command with the given arguments.
// It connects standard output and error to the current process for visibility.
func (g *GitRunner) run(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

// runAll runs all necessary git commands: add, commit, and push.
func (g *GitRunner) runAll(commitMessage string) error {
	if err := g.run("git", "add", "."); err != nil {
		return fmt.Errorf("git add failed: %w", err)
	}

	if err := g.run("git", "commit", "-m", commitMessage); err != nil {
		return fmt.Errorf("git commit failed: %w", err)
	}

	if err := g.run("git", "push", "origin", "main"); err != nil {
		return fmt.Errorf("git push failed: %w", err)
	}

	return nil
}

// main is the entry point of the automation pod.
func main() {
	// Initialize the generator instance
	gen := generator.NewGenerator()

	// Generate a new file
	if err := gen.GenerateFile(); err != nil {
		log.Fatalf("error generating file: %v", err)
	}

	// Initialize Git runner
	gr := &GitRunner{}

	// Run all Git commands (add, commit, push)
	if err := gr.runAll("Automated file generation"); err != nil {
		log.Fatalf("git operation failed: %v", err)
	}

	fmt.Println("✅ File generated and pushed successfully.")
}
