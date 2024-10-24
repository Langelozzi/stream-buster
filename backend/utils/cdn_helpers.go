package utils

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

// DoesContentExist Asynchronous function to check if content exists for the given TMDB ID
func DoesContentExist(tmdbId int, isTV bool) bool {
	baseUrl := GetEnvVariable("VIDSRC_BASE_URL")

	var url string
	if isTV {
		url = fmt.Sprintf("%s/tv/%s", baseUrl, strconv.Itoa(tmdbId))
	} else {
		url = fmt.Sprintf("%s/movie/%s", baseUrl, strconv.Itoa(tmdbId))
	}

	// Ping the vidsrc API to check if the content exists
	response, err := http.Get(url)
	//body, err := io.ReadAll(response.Body)
	if err != nil || response.StatusCode != http.StatusOK {
		fmt.Println(response.StatusCode)
		return false
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(response.Body)

	fmt.Println(response.StatusCode)
	return true
}
