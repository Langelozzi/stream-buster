package utils

import (
	"io"
	"log"
	"net/http"
)

func GetAsync(url string) (string, error) {
	response, err := http.Get(url)
	if err != nil {
		log.Fatalf("Error making GET request: %v\n", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			return
		}
	}(response.Body)

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v\n", err)
	}

	return string(body), nil
}
