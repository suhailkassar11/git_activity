package api

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/suhailkassar11/github_activity/internal/models"
)

// FetchUser fetches GitHub events for a given username and returns the API response.
func FetchUser(username string) (models.API_Response, error) {
	// Initialize an empty API_Response to return in case of error
	var rawEvents models.API_Response

	// Construct the GitHub API URL
	baseURL := "https://api.github.com"
	url := fmt.Sprintf("%s/users/%s/events", baseURL, username)

	// Create a new HTTP request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return rawEvents, fmt.Errorf("error creating request: %w", err)
	}

	// Set headers (optional, but recommended for GitHub API)
	req.Header.Set("Accept", "application/vnd.github.v3+json")

	// Create an HTTP client and execute the request
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return rawEvents, fmt.Errorf("error fetching data: %w", err)
	}
	defer resp.Body.Close()

	// Check the response status code
	if resp.StatusCode != http.StatusOK {
		return rawEvents, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	// Decode the response body into rawEvents
	if err := json.NewDecoder(resp.Body).Decode(&rawEvents); err != nil {
		return rawEvents, fmt.Errorf("error decoding JSON: %w", err)
	}

	return rawEvents, nil
}
