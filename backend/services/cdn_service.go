package services

import (
	iDao "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"net/http"
)

type CDNService struct {
	dao iDao.CDNDaoInterface
}

func NewCDNService(dao iDao.CDNDaoInterface) *CDNService {
	return &CDNService{dao: dao}
}

func (service *CDNService) GetMovieContent(tmdbId string) (string, error) {
	html, err := service.dao.GetMovieContent(tmdbId)
	if err != nil {
		return "", err
	}
	return html, nil
}

func (service *CDNService) GetTVContent(tmdbId string, seasonNum int, episodeNum int) (string, error) {
	html, err := service.dao.GetTVContent(tmdbId, seasonNum, episodeNum)
	if err != nil {
		return "", err
	}
	return html, nil
}

func (service *CDNService) CheckContentExists(tmdbId string, isTV bool) int {
	// Exists = 1
	// Doesn't exist = 0
	// Requests exceed or error occurred (i.e. unknown) = -1
	res, err := service.dao.CheckContentExist(tmdbId, isTV)
	if err != nil {
		return -1
	}

	if res.StatusCode == http.StatusOK {
		return 1
	} else if res.StatusCode == http.StatusNotFound {
		return 0
	}

	return -1
}
