package githubutils

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

type GitHubEvent struct {
	Type string `json:"type"`

	Repo struct {
		Name string
	} `json:"repo"`

	CreatedAt string `json:"created_at"`
}

func (event GitHubEvent) Print() {
	fmt.Printf("Type: %s\n", event.Type)
	fmt.Printf("Created At: %s\n", event.CreatedAt)
	fmt.Printf("Reprository: %s\n", event.Repo.Name)
}

func GetGitHubData(username string, url string) ([]GitHubEvent, error) {
	furl := fmt.Sprintf(url, username)

	// Fetch the api
	resp, err := http.Get(furl)
	if err != nil {
		return []GitHubEvent{}, err
	}

	// Read data
	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return []GitHubEvent{}, err
	}

	var dat []GitHubEvent

	if err := json.Unmarshal(data, &dat); err != nil {
		return []GitHubEvent{}, err
	}

	return dat, nil
}
