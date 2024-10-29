package daos

import (
	"fmt"
	"github.com/STREAM-BUSTER/stream-buster/adapters"
	"github.com/STREAM-BUSTER/stream-buster/models/api"
	"github.com/STREAM-BUSTER/stream-buster/utils"
	"io"
	"net/http"
	"net/url"
	"strconv"
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

	var allResults []interface{}
	page := 1

	for {
		getUrl := fmt.Sprintf("%s%s?api_key=%s&query=%s&page=%d", baseUrl, relativeUrl, apiKey, encodedQuery, page)

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

		// Get how many pages there are
		totalPages, err := adapters.GetTotalPageCount(string(body))
		if err != nil {
			return nil, err
		}

		// Parse the response into a structure that includes page and total_pages
		parsedResponse, err := adapters.ParseSearchMultiMediaResponse(string(body))
		if err != nil {
			return nil, err
		}

		// Append the results
		allResults = append(allResults, parsedResponse...)

		// Check if we've fetched all pages
		if page >= totalPages {
			break
		}
		page++
	}

	return allResults, nil

}

func (dao *TMDBDao) GetTVDetails(id int) (*api.TV, error) {
	baseUrl := utils.GetEnvVariable("TMDB_API_BASE_URL")
	apiKey := utils.GetEnvVariable("TMDB_API_KEY")
	relativeUrl := "/tv"

	getUrl := fmt.Sprintf("%s%s/%s?api_key=%s", baseUrl, relativeUrl, strconv.Itoa(id), apiKey)

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

	fmt.Println(body)

	tvObj, err := adapters.ParseTVDetailsResponse(string(body))
	if err != nil {
		fmt.Printf("Error casting response to TV object: %v\n", err)
		return nil, err
	}

	return tvObj, nil
}
