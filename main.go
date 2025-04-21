package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"time"

	"github.com/Phani0310/github-automation-pod/generator"
	"github.com/Phani0310/github-automation-pod/internal"
)

func run(command string, args ...string) error {
	cmd := exec.Command(command, args...)
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	return cmd.Run()
}

func main() {
	token := os.Getenv("GITHUB_TOKEN")
	if token == "" {
		log.Fatal("GITHUB_TOKEN is not set")
	}

	watcher := internal.NewChecker("Phani0310/github-automation-pod", "main", token)
	watcher.LoadState()

	for {
		newCommit, err := watcher.IsNewCommit()
		if err != nil {
			log.Println("check failed:", err)
			time.Sleep(30 * time.Second)
			continue
		}

		if newCommit {
			fmt.Println("New commit detected! Generating file...")

			gen := generator.New()
			if _, err := gen.Generate(); err != nil {
				log.Println("generation failed:", err)
				continue
			}

			if err := run("git", "add", "."); err != nil {
				log.Fatal("git add failed:", err)
			}
			if err := run("git", "commit", "-m", "Auto: generated file after new commit"); err != nil {
				log.Println("git commit skipped (maybe no changes):", err)
			}
			if err := run("git", "push"); err != nil {
				log.Fatal("git push failed:", err)
			}

			watcher.SaveState()
		}

		time.Sleep(30 * time.Second)
	}
}
