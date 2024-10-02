package main

import (
	"fmt"
	"net/url"
	"strings"
)

func normalizeURL(inputURL string) (string, error) {
	parsedURL, err := url.Parse(inputURL)
	if err != nil {
		return "", fmt.Errorf("couldn't parse URL: %w", err)
	}

	normalizedURL := parsedURL.Host + parsedURL.Path
	normalizedURL = strings.ToLower(normalizedURL)
	normalizedURL = strings.TrimRight(normalizedURL, "/")
	return normalizedURL, nil
}
