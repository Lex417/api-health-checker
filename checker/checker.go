package checker

import (
	"fmt"
	"net/http"
	"time"
)

func CheckURL(url string) (int, error) {
	client := http.Client{
		Timeout: 5 * time.Second,
	}

	resp, err := client.Get(url)
	if err != nil {
		return 0, fmt.Errorf("HTTP GET request failed for %s: %w", url, err)
	}
	defer resp.Body.Close()

	return resp.StatusCode, nil
}
