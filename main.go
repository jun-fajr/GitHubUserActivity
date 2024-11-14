package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// GitHubEvent represents a simplified structure for GitHub events
type GitHubEvent struct {
	Type string `json:"type"`
	Repo struct {
		Name string `json:"name"`
	} `json:"repo"`
}

// getUserActivity fetches the recent activity for a specified GitHub user
func getUserActivity(username string) ([]GitHubEvent, error) {
	url := fmt.Sprintf("https://api.github.com/users/%s/events", username)
	resp, err := http.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch data: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("error: received status %s", resp.Status)
	}

	var events []GitHubEvent
	if err := json.NewDecoder(resp.Body).Decode(&events); err != nil {
		return nil, fmt.Errorf("failed to parse JSON: %v", err)
	}

	return events, nil
}

// displayActivity processes and prints the GitHub user activity
func displayActivity(events []GitHubEvent) {
	for _, event := range events {
		var action string
		switch event.Type {
		case "PushEvent":
			action = "Pushed to"
		case "IssuesEvent":
			action = "Created an issue in"
		case "WatchEvent":
			action = "Starred"
		case "ForkEvent":
			action = "Forked"
		case "CreateEvent":
			action = "Created a repository"
		default:
			action = event.Type
		}
		fmt.Printf("- %s %s\n", action, event.Repo.Name)
	}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: github-activity <username>")
		os.Exit(1)
	}

	username := os.Args[1]

	fmt.Printf("Fetching activity for user: %s\n", username)
	events, err := getUserActivity(username)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}

	if len(events) == 0 {
		fmt.Println("No recent activity found.")
		return
	}

	displayActivity(events)
}
