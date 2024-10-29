package daos

import (
	"fmt"
	"github.com/STREAM-BUSTER/stream-buster/utils"
	"io"
	"log"
	"net/http"
	"strconv"
)

type CDNDao struct{}

func NewCDNDao() *CDNDao {
	return &CDNDao{}
}

func (dao *CDNDao) GetMovieContent(tmdbId string) (string, error) {
	baseUrl := utils.GetEnvVariable("VIDSRC_BASE_URL")
	url := fmt.Sprintf("%s/movie/%s", baseUrl, tmdbId)

	// Make the first GET request
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

func (dao *CDNDao) GetTVContent(tmdbId string, seasonNum int, episodeNum int) (string, error) {
	// TODO: make util function for building the vidsrc url
	baseUrl := utils.GetEnvVariable("VIDSRC_BASE_URL")
	url := fmt.Sprintf("%s/tv/%s/%s-%s", baseUrl, tmdbId, strconv.Itoa(seasonNum), strconv.Itoa(episodeNum))

	// Make the first GET request
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
