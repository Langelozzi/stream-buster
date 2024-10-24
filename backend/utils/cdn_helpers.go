package utils

import (
	"fmt"
	"net/http"
	"strconv"
)

func DoesContentExistForTMDBId(tmdbId int) bool {
	baseUrl := GetEnvVariable("VIDSRC_BASE_URL")
	url := fmt.Sprintf("%s/movie/%s", baseUrl, strconv.Itoa(tmdbId))

	// Ping the vidsrc api to check if the content exists
	response, err := http.Get(url)
	if err != nil {
		return false
	}

	if response.StatusCode != http.StatusOK {
		return false
	}
	return true
}
