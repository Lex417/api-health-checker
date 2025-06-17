package checker

import (
	"fmt"
	"net/http"
	"time"
)

// CheckURL performs an HTTP GET request to the given URL and returns its status code and any error.
func CheckURL(url string) (int, error) {
	// Create a custom HTTP client with a timeout to prevent hanging requests.
	client := http.Client{
		Timeout: 5 * time.Second, // Set a timeout for the HTTP request.
	}

	// Perform the GET request.
	resp, err := client.Get(url)
	if err != nil {
		// If there's a network error or DNS resolution error, return 0 status and the error.
		return 0, fmt.Errorf("HTTP GET request failed for %s: %w", url, err)
	}
	defer resp.Body.Close() // Ensure the response body is closed to prevent resource leaks.

	// Return the HTTP status code.
	return resp.StatusCode, nil
}
