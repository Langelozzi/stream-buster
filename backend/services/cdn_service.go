package services

import (
	"errors"
	iDao "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"regexp"
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
	return transformContent(html)
}

func (service *CDNService) GetTVContent(tmdbId string, seasonNum int, episodeNum int) (string, error) {
	html, err := service.dao.GetTVContent(tmdbId, seasonNum, episodeNum)
	if err != nil {
		return "", nil
	}
	return transformContent(html)
}

func (service *CDNService) CheckContentExists(tmdbId string, isTV bool) bool {
	return service.dao.CheckContentExist(tmdbId, isTV)
}

func transformContent(html string) (string, error) {
	srcUrl, err := getContentSrcUrl(html)
	if err != nil {
		return "", err
	}

	wrappedHtml := getWrappedHtmlContent(srcUrl)

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

func getWrappedHtmlContent(contentSrcUrl string) string {
	// Create a wrapped HTML with an iframe
	wrappedHTML := `
		<!DOCTYPE html>
		<html lang="en">
		<head>
			<meta charset="UTF-8">
			<title>Video Embed</title>
			<style>
				iframe {
					width: 100%;
					height: 100vh; /* Use viewport height for the iframe */
					border: none;
				}
				html, body {
					margin: 0;
					padding: 0;
					width: 100%;
					height: 100%;
					overflow: hidden; /* Hide overflow to prevent scrollbars */
				}
			</style>
		</head>
		<body>
			<iframe src="` + contentSrcUrl + `" allowFullScreen sandbox="allow-forms allow-pointer-lock allow-same-origin allow-scripts allow-top-navigation"></iframe>
		</body>
		</html>`

	return wrappedHTML
}
