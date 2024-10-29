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
					tv := CastToTV(itemMap, media, false)
					castedResults = append(castedResults, tv)
				}
			}
		}
	} else {
		fmt.Println("No results found or results is not an array")
	}

	return castedResults, nil
}

func ParseTVDetailsResponse(json string) (*api.TV, error) {
	jsonMap, err := JSONToMap(json)
	if err != nil {
		return nil, err
	}

	mediaType := &db.MediaType{
		Name: "tv",
	}
	media := CastToMedia(jsonMap, mediaType)
	tv := CastToTV(jsonMap, media, true)

	return tv, nil
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
	overview := obj["overview"].(string)

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
		Overview:    overview,
		PosterImage: posterPath,
		MediaType:   mediaType,
	}
}

func CastToMovie(obj map[string]interface{}, media *db.Media) *api.Movie {
	releaseDate := ConvertStringToDate(obj["release_date"].(string))

	return &api.Movie{
		Media:       media,
		ReleaseDate: releaseDate,
	}
}

func CastToTV(obj map[string]interface{}, media *db.Media, detailed bool) *api.TV {
	firstAirDate := ConvertStringToDate(obj["first_air_date"].(string))

	tv := api.TV{
		Media:        media,
		FirstAirDate: firstAirDate,
	}

	if detailed {
		tv.LastAirDate = ConvertStringToDate(obj["last_air_date"].(string))
		tv.SeasonCount = int(obj["number_of_seasons"].(float64))
		tv.EpisodeCount = int(obj["number_of_episodes"].(float64))
		tv.BackdropImage = GetFullImagePath(obj["backdrop_path"].(string))

		// Iterate through the seasons list and convert to season objects
		var castedSeasons []*api.Season
		if seasons, ok := obj["seasons"].([]interface{}); ok {
			for _, season := range seasons {
				if seasonMap, ok := season.(map[string]interface{}); ok {
					casted := CastToSeason(seasonMap, media)
					castedSeasons = append(castedSeasons, casted)
				}
			}
		}

		tv.Seasons = castedSeasons
	}

	return &tv
}

func CastToSeason(obj map[string]interface{}, media *db.Media) *api.Season {
	tmdbId := int(obj["id"].(float64))
	seasonNum := int(obj["season_number"].(float64))
	episodeCount := int(obj["episode_count"].(float64))
	name := obj["name"].(string)
	overview := obj["overview"].(string)
	posterPath := GetFullImagePath(obj["poster_path"].(string))

	season := api.Season{
		Media:        media,
		SeasonTMDBID: tmdbId,
		SeasonNumber: seasonNum,
		EpisodeCount: episodeCount,
		Name:         name,
		Overview:     overview,
		PosterPath:   posterPath,
	}

	return &season
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
