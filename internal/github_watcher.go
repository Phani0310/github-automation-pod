package internal

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// CommitChecker polls GitHub for the latest commit on a branch
type CommitChecker struct {
	Repo      string // e.g. "Phani0310/github-automation-pod"
	Branch    string // typically "main"
	LastSHA   string
	Token     string
	StateFile string
}

// NewChecker initializes with repo and token
func NewChecker(repo, branch, token string) *CommitChecker {
	return &CommitChecker{
		Repo:      repo,
		Branch:    branch,
		Token:     token,
		StateFile: ".last_sha",
	}
}

// LoadState loads the last known SHA
func (c *CommitChecker) LoadState() {
	data, err := os.ReadFile(c.StateFile)
	if err == nil {
		c.LastSHA = string(data)
	}
}

// SaveState saves the current SHA to disk
func (c *CommitChecker) SaveState() {
	_ = os.WriteFile(c.StateFile, []byte(c.LastSHA), 0644)
}

// IsNewCommit checks GitHub for new commits
func (c *CommitChecker) IsNewCommit() (bool, error) {
	url := fmt.Sprintf("https://api.github.com/repos/%s/commits/%s", c.Repo, c.Branch)
	req, _ := http.NewRequest("GET", url, nil)
	req.Header.Add("Authorization", "Bearer "+c.Token)
	req.Header.Add("Accept", "application/vnd.github+json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return false, err
	}
	defer res.Body.Close()

	var response struct {
		SHA string `json:"sha"`
	}
	if err := json.NewDecoder(res.Body).Decode(&response); err != nil {
		return false, err
	}

	if response.SHA != "" && response.SHA != c.LastSHA {
		c.LastSHA = response.SHA
		return true, nil
	}
	return false, nil
}
