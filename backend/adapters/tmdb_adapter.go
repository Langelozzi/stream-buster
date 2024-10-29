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
					movie := CastToMovie(itemMap, media, false)
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

func ParseMovieDetailsResponse(json string) (*api.Movie, error) {
	jsonMap, err := JSONToMap(json)
	if err != nil {
		return nil, err
	}

	mediaType := &db.MediaType{
		Name: "movie",
	}
	media := CastToMedia(jsonMap, mediaType)
	movie := CastToMovie(jsonMap, media, true)

	return movie, nil
}

func ParseEpisodeListResponse(json string) ([]*api.Episode, error) {
	jsonMap, err := JSONToMap(json)
	if err != nil {
		return nil, err
	}

	var castedResults []*api.Episode
	// Access the "results" array
	if results, ok := jsonMap["episodes"].([]interface{}); ok {
		// Iterate through the results and cast them to our structs
		for _, item := range results {
			if itemMap, ok := item.(map[string]interface{}); ok {
				episode := CastToEpisode(itemMap, nil)
				castedResults = append(castedResults, episode)
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
	var mediaType string
	if obj["media_type"] != nil {
		if mt, ok := obj["media_type"].(string); ok {
			mediaType = mt
		}
	}

	return &db.MediaType{
		Name: mediaType,
	}
}

func CastToMedia(obj map[string]interface{}, mediaType *db.MediaType) *db.Media {
	var tmdbId int
	if obj["id"] != nil {
		if id, ok := obj["id"].(float64); ok {
			tmdbId = int(id)
		}
	}

	var overview string
	if obj["overview"] != nil {
		overview = obj["overview"].(string)
	}

	var title string
	if mediaType.Name == "tv" {
		if obj["name"] != nil {
			title = obj["name"].(string)
		}
	} else {
		if obj["title"] != nil {
			title = obj["title"].(string)
		}
	}

	var posterPath string
	if obj["poster_path"] != nil {
		if path, ok := obj["poster_path"].(string); ok {
			posterPath = GetFullImagePath(path)
		}
	}

	var castedGenres []*db.Genre
	if obj["genres"] != nil {
		if genreMaps, ok := obj["genres"].([]interface{}); ok {
			for _, genreMap := range genreMaps {
				if genre, ok := genreMap.(map[string]interface{}); ok {
					castedGenre := CastToGenre(genre)
					castedGenres = append(castedGenres, castedGenre)
				}
			}
		}
	}

	return &db.Media{
		TMDBID:      tmdbId,
		Title:       title,
		Overview:    overview,
		PosterImage: posterPath,
		MediaType:   mediaType,
		Genres:      castedGenres,
	}
}

func CastToMovie(obj map[string]interface{}, media *db.Media, detailed bool) *api.Movie {
	var releaseDate *time.Time
	if obj["release_date"] != nil {
		if dateStr, ok := obj["release_date"].(string); ok {
			releaseDate = ConvertStringToDate(dateStr)
		}
	}

	movie := api.Movie{
		Media:       media,
		ReleaseDate: releaseDate,
	}

	if detailed {
		if obj["backdrop_path"] != nil {
			if backdropPath, ok := obj["backdrop_path"].(string); ok {
				movie.BackdropImage = GetFullImagePath(backdropPath)
			}
		}

		if obj["runtime"] != nil {
			if runtime, ok := obj["runtime"].(float64); ok {
				movie.Runtime = int(runtime)
			}
		}
	}

	return &movie
}

func CastToTV(obj map[string]interface{}, media *db.Media, detailed bool) *api.TV {
	var firstAirDate *time.Time
	if obj["first_air_date"] != nil {
		if dateStr, ok := obj["first_air_date"].(string); ok {
			firstAirDate = ConvertStringToDate(dateStr)
		}
	}

	tv := api.TV{
		Media:        media,
		FirstAirDate: firstAirDate,
	}

	if detailed {
		if obj["last_air_date"] != nil {
			if dateStr, ok := obj["last_air_date"].(string); ok {
				tv.LastAirDate = ConvertStringToDate(dateStr)
			}
		}

		if obj["number_of_seasons"] != nil {
			if seasonCount, ok := obj["number_of_seasons"].(float64); ok {
				tv.SeasonCount = int(seasonCount)
			}
		}

		if obj["number_of_episodes"] != nil {
			if episodeCount, ok := obj["number_of_episodes"].(float64); ok {
				tv.EpisodeCount = int(episodeCount)
			}
		}

		if obj["backdrop_path"] != nil {
			if backdropPath, ok := obj["backdrop_path"].(string); ok {
				tv.BackdropImage = GetFullImagePath(backdropPath)
			}
		}

		var castedSeasons []*api.Season
		if obj["seasons"] != nil {
			if seasons, ok := obj["seasons"].([]interface{}); ok {
				for _, season := range seasons {
					if seasonMap, ok := season.(map[string]interface{}); ok {
						casted := CastToSeason(seasonMap, media)
						castedSeasons = append(castedSeasons, casted)
					}
				}
			}
		}
		tv.Seasons = castedSeasons
	}

	return &tv
}

func CastToSeason(obj map[string]interface{}, media *db.Media) *api.Season {
	var tmdbId, seasonNum, episodeCount int
	if obj["id"] != nil {
		if id, ok := obj["id"].(float64); ok {
			tmdbId = int(id)
		}
	}
	if obj["season_number"] != nil {
		if sn, ok := obj["season_number"].(float64); ok {
			seasonNum = int(sn)
		}
	}
	if obj["episode_count"] != nil {
		if ec, ok := obj["episode_count"].(float64); ok {
			episodeCount = int(ec)
		}
	}

	var name, overview, posterPath string
	if obj["name"] != nil {
		if n, ok := obj["name"].(string); ok {
			name = n
		}
	}
	if obj["overview"] != nil {
		if ov, ok := obj["overview"].(string); ok {
			overview = ov
		}
	}
	if obj["poster_path"] != nil {
		if pp, ok := obj["poster_path"].(string); ok {
			posterPath = GetFullImagePath(pp)
		}
	}

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

func CastToEpisode(obj map[string]interface{}, media *db.Media) *api.Episode {
	var name string
	var overview string
	var id int
	var episodeNum int
	var stillPath string
	var runtime int
	var seasonNum int

	if obj["id"] != nil {
		if idFloat, ok := obj["id"].(float64); ok {
			id = int(idFloat)
		}
	}

	if obj["name"] != nil {
		if nameStr, ok := obj["name"].(string); ok {
			name = nameStr
		}
	}

	if obj["overview"] != nil {
		if overviewStr, ok := obj["overview"].(string); ok {
			overview = overviewStr
		}
	}

	if obj["episode_number"] != nil {
		if episodeNumFloat, ok := obj["episode_number"].(float64); ok {
			episodeNum = int(episodeNumFloat)
		}
	}

	if obj["still_path"] != nil {
		if stillPathStr, ok := obj["still_path"].(string); ok {
			stillPath = GetFullImagePath(stillPathStr)
		}
	}

	if obj["runtime"] != nil {
		if runtimeFloat, ok := obj["runtime"].(float64); ok {
			runtime = int(runtimeFloat)
		}
	}

	if obj["season_number"] != nil {
		if seasonNumFloat, ok := obj["season_number"].(float64); ok {
			seasonNum = int(seasonNumFloat)
		}
	}

	episode := api.Episode{
		Name:          name,
		Overview:      overview,
		EpisodeTMDBID: id,
		EpisodeNumber: episodeNum,
		StillPath:     stillPath,
		Runtime:       runtime,
		SeasonNumber:  seasonNum,
	}

	if media != nil {
		episode.Media = media
	}

	return &episode
}

func CastToGenre(obj map[string]interface{}) *db.Genre {
	var id int
	var name string

	if obj["id"] != nil {
		if idFloat, ok := obj["id"].(float64); ok {
			id = int(idFloat)
		}
	}

	if obj["name"] != nil {
		if nameStr, ok := obj["name"].(string); ok {
			name = nameStr
		}
	}

	return &db.Genre{
		ID:   id,
		Name: name,
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
