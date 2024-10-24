package adapters

import (
	"encoding/json"
	"fmt"
	"github.com/STREAM-BUSTER/stream-buster/models/api"
	"github.com/STREAM-BUSTER/stream-buster/models/db"
	"github.com/STREAM-BUSTER/stream-buster/utils"
	"time"
)

func ParseSearchMultiMediaResponse(json string) ([]interface{}, error) {
	jsonMap, err := JSONToMap(json)
	if err != nil {
		return nil, err
	}

	var castedResults []interface{}
	// Access the "results" array
	if results, ok := jsonMap["results"].([]interface{}); ok {
		// Iterate through the results and cast them to our structs
		for _, item := range results {
			if itemMap, ok := item.(map[string]interface{}); ok {
				mediaType := CastToMediaType(itemMap)
				if mediaType.Name != "tv" && mediaType.Name != "movie" {
					continue
				}
				media := CastToMedia(itemMap, mediaType)

				if mediaType.Name == "movie" {
					movie := CastToMovie(itemMap, media)
					castedResults = append(castedResults, movie)
				} else if mediaType.Name == "tv" {
					tv := CastToTV(itemMap, media)
					castedResults = append(castedResults, tv)
				}
			}
		}
	} else {
		fmt.Println("No results found or results is not an array")
	}

	return castedResults, nil
}

func GetTotalPageCount(json string) (int, error) {
	jsonMap, err := JSONToMap(json)
	if err != nil {
		return -1, err
	}

	return int(jsonMap["total_pages"].(float64)), nil
}

func CastToMediaType(obj map[string]interface{}) *db.MediaType {
	mediaType := ""
	if obj["media_type"] != nil {
		mediaType = obj["media_type"].(string)
	}

	return &db.MediaType{
		Name: mediaType,
	}
}

func CastToMedia(obj map[string]interface{}, mediaType *db.MediaType) *db.Media {
	tmdbId := int(obj["id"].(float64))

	var title string
	if mediaType.Name == "tv" {
		title = obj["name"].(string)
	} else {
		title = obj["title"].(string)
	}

	posterPath := ""
	if obj["poster_path"] != nil {
		posterPath = GetFullImagePath(obj["poster_path"].(string))
	}

	return &db.Media{
		TMDBID:      tmdbId,
		Title:       title,
		PosterImage: posterPath,
		MediaType:   mediaType,
	}
}

func CastToMovie(obj map[string]interface{}, media *db.Media) *api.Movie {
	overview := obj["overview"].(string)
	releaseDate := ConvertStringToDate(obj["release_date"].(string))

	return &api.Movie{
		Media:       media,
		Overview:    overview,
		ReleaseDate: releaseDate,
	}
}

func CastToTV(obj map[string]interface{}, media *db.Media) *api.TV {
	overview := obj["overview"].(string)
	firstAirDate := ConvertStringToDate(obj["first_air_date"].(string))

	return &api.TV{
		Media:        media,
		Overview:     overview,
		FirstAirDate: firstAirDate,
	}
}

func ConvertStringToDate(str string) *time.Time {
	layout := "2006-01-02"

	parsedDate, err := time.Parse(layout, str)
	if err != nil {
		fmt.Println("Error parsing date:", err)
		return nil
	}

	return &parsedDate
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
