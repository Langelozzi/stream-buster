package adapters

import (
	"encoding/json"
	"fmt"
	"github.com/STREAM-BUSTER/stream-buster/utils"
)

func ParseSearchMultiMediaResponse(json string) ([]interface{}, error) {
	jsonMap, err := JSONToMap(json)
	if err != nil {
		return nil, err
	}

	// Access the "results" array
	if results, ok := jsonMap["results"].([]interface{}); ok {
		// Iterate through the results and print the information
		for _, item := range results {
			if itemMap, ok := item.(map[string]interface{}); ok {
				// Media object info
				tmdbId := itemMap["id"].(int)
				title := itemMap["name"].(string)
				posterPath := GetFullImagePath(itemMap["poster_path"].(string))

				// Media type object info
				mediaType := itemMap["media_type"].(string)

				// Movie object info
				overview := itemMap["overview"].(string)
				releaseDate := itemMap["release_date"].(string)

				// Ideas:
				// Make helper functions for creating each individual object, media, media_type, movie, tv
				// The movie one will require a media object to be passed in, etc.
				// ...
				// Make a conditional here for if mediatype is tv or movie

			}
		}
	} else {
		fmt.Println("No results found or results is not an array")
	}

	fmt.Println(jsonMap)

	return []interface{}{}, nil
}

func GetFullImagePath(relativePath string) string {
	baseUrl := utils.GetEnvVariable("TMDB_IMAGE_BASE_URL")
	fullUrl := baseUrl + relativePath

	return fullUrl
}

// JSONToMap converts a JSON string to a nested map.
func JSONToMap(jsonStr string) (map[string]interface{}, error) {
	var result map[string]interface{}
	err := json.Unmarshal([]byte(jsonStr), &result)
	if err != nil {
		return nil, fmt.Errorf("error unmarshaling JSON: %v", err)
	}
	return result, nil
}
