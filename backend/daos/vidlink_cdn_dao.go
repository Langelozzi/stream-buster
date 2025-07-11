package daos

import (
	"fmt"
	"github.com/STREAM-BUSTER/stream-buster/utils"
	"net/http"
	"strconv"
)

type VidLinkDao struct{}

func NewVidLinkDao() *VidLinkDao {
	return &VidLinkDao{}
}

func (dao *VidLinkDao) GetMovieContent(tmdbId string) (string, error) {
	baseUrl := utils.GetEnvVariable("VIDLINK_BASE_URL")
	url := fmt.Sprintf("%s/movie/%s", baseUrl, tmdbId)
	return utils.GetWrappedHtmlContent(url), nil
}

func (dao *VidLinkDao) GetTVContent(tmdbId string, seasonNum int, episodeNum int) (string, error) {
	baseUrl := utils.GetEnvVariable("VIDLINK_BASE_URL")
	url := fmt.Sprintf("%s/tv/%s/%s/%s", baseUrl, tmdbId, strconv.Itoa(seasonNum), strconv.Itoa(episodeNum))
	return utils.GetWrappedHtmlContent(url), nil
}

// CheckContentExist Asynchronous function to check if content exists for the given TMDB ID
func (dao *VidLinkDao) CheckContentExist(tmdbId string, isTV bool) (*http.Response, error) {
	baseUrl := utils.GetEnvVariable("VIDLINK_BASE_URL")

	var url string
	if isTV {
		url = fmt.Sprintf("%s/tv/%s/1/1", baseUrl, tmdbId)
	} else {
		url = fmt.Sprintf("%s/movie/%s", baseUrl, tmdbId)
	}

	// Ping the API to check if the content exists
	return http.Get(url)
}
