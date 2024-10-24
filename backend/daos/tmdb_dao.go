package daos

import (
	"fmt"
	"github.com/STREAM-BUSTER/stream-buster/adapters"
	"github.com/STREAM-BUSTER/stream-buster/utils"
	"io"
	"net/http"
	"net/url"
)

type TMDBDao struct{}

func NewTMDBDao() *TMDBDao {
	return &TMDBDao{}
}

func (dao *TMDBDao) SearchMultiMedia(query string) ([]interface{}, error) {
	// Get environment variables
	baseUrl := utils.GetEnvVariable("TMDB_API_BASE_URL")
	apiKey := utils.GetEnvVariable("TMDB_API_KEY")
	relativeUrl := "/search/multi"
	encodedQuery := url.QueryEscape(query)

	getUrl := fmt.Sprintf("%s%s?api_key=%s&query=%s", baseUrl, relativeUrl, apiKey, encodedQuery)

	response, err := http.Get(getUrl)
	if err != nil {
		fmt.Printf("Error fetching from tmdb api: %v\n", err)
		return nil, err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Printf("Error closing request body stream: %v\n", err)
		}
	}(response.Body)

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("Error reading response body: %v\n", err)
		return nil, err
	}

	fmt.Printf("Response body: %v\n", string(body))

	return adapters.ParseSearchMultiMediaResponse(string(body))
}
