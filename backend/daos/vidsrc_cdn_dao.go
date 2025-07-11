package daos

import (
	"errors"
	"fmt"
	"github.com/STREAM-BUSTER/stream-buster/utils"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type VidSrcDao struct{}

func NewVidSrcDao() *VidSrcDao {
	return &VidSrcDao{}
}

func (dao *VidSrcDao) GetMovieContent(tmdbId string) (string, error) {
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

	return transformContent(string(body))
}

func (dao *VidSrcDao) GetTVContent(tmdbId string, seasonNum int, episodeNum int) (string, error) {
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

	return transformContent(string(body))
}

// CheckContentExist Asynchronous function to check if content exists for the given TMDB ID
func (dao *VidSrcDao) CheckContentExist(tmdbId string, isTV bool) (*http.Response, error) {
	baseUrl := utils.GetEnvVariable("VIDSRC_BASE_URL")

	var url string
	if isTV {
		url = fmt.Sprintf("%s/tv/%s", baseUrl, tmdbId)
	} else {
		url = fmt.Sprintf("%s/movie/%s", baseUrl, tmdbId)
	}

	// Ping the vidsrc API to check if the content exists
	return http.Get(url)
}

func transformContent(html string) (string, error) {
	srcUrl, err := getContentSrcUrl(html)
	if err != nil {
		return "", err
	}

	innerContentBody, _ := utils.GetAsync(srcUrl)
	println(innerContentBody)

	wrappedHtml := utils.GetWrappedHtmlContent(srcUrl)

	return wrappedHtml, nil
}

func getContentSrcUrl(html string) (string, error) {
	// Use a regex to find the src attribute of the iframe
	re := regexp.MustCompile(`src="([^"]+)"`)
	match := re.FindStringSubmatch(html)

	if len(match) == 0 {
		return "", errors.New("no src attribute found")
	}

	// Extract the src value and construct the full URL
	src := match[1]
	srcUrl := "https:" + src

	return srcUrl, nil
}
