package controllers

import (
	"fmt"
	"github.com/STREAM-BUSTER/stream-buster/services/interfaces"
	"github.com/gin-gonic/gin"
	"io"
	"log"
	"net/http"
	"regexp"
	"strconv"
)

type CDNController struct {
	service interfaces.CDNServiceInterface
}

func NewCDNController(service interfaces.CDNServiceInterface) *CDNController {
	return &CDNController{
		service: service,
	}
}

// GetMultiMediaSearchResults retrieves multimedia search results based on a query.
// @Summary Retrieve multimedia search results
// @Description Get multimedia content based on the search query.
// @Tags media
// @Accept  json
// @Produce  json
// @Param query query string true "Search query for multimedia content"
// @Success 200 {object} []interface{} "Successfully retrieved multimedia search results"
// @Failure 400 {object} map[string]interface{} "Error: Invalid or empty query, or no results found"
// @Router /media/search [get]
func (contr *CDNController) GetMovieContent(c *gin.Context) {
	// get the tmdbId from the route params
	tmdbId := c.Param("tmdbId")

	// call the service
	html, err := contr.service.GetMovieContent(tmdbId)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Unable to procure content for id: " + tmdbId,
			"error":   err,
		})
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}

// GetMultiMediaSearchResults retrieves multimedia search results based on a query.
// @Summary Retrieve multimedia search results
// @Description Get multimedia content based on the search query.
// @Tags media
// @Accept  json
// @Produce  json
// @Param query query string true "Search query for multimedia content"
// @Success 200 {object} []interface{} "Successfully retrieved multimedia search results"
// @Failure 400 {object} map[string]interface{} "Error: Invalid or empty query, or no results found"
// @Router /media/search [get]
func (contr *CDNController) GetTVContent(c *gin.Context) {
	// get the tmdbId from the route params
	tmdbId := c.Param("tmdbId")

	// get the season num and convert to int
	seasonNum, err := strconv.Atoi(c.Param("seasonNum"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "seasonNum must be an integer",
		})
		return
	}

	// get the episode num and convert to int
	episodeNum, err := strconv.Atoi(c.Param("episodeNum"))
	if err != nil {
		c.JSON(400, gin.H{
			"message": "episodeNum must be an integer",
		})
		return
	}

	// call the service
	html, err := contr.service.GetTVContent(tmdbId, seasonNum, episodeNum)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Unable to procure content for id: " + tmdbId,
			"error":   err,
		})
		return
	}

	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(html))
}

func (contr *CDNController) TestContentProxy(c *gin.Context) {
	vidSrcUrl := "https://vidsrc.xyz/embed/movie/tt5433140"

	// Make the first GET request
	response, err := http.Get(vidSrcUrl)
	if err != nil {
		log.Fatalf("Error making GET request: %v\n", err)
	}
	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v\n", err)
	}

	// Use a regex to find the src attribute of the iframe
	re := regexp.MustCompile(`src="([^"]+)"`)
	match := re.FindStringSubmatch(string(body))

	if len(match) == 0 {
		fmt.Println("No src attribute found")
		c.String(http.StatusNotFound, "No src attribute found")
		return
	}

	// Extract the src value and construct the full URL
	src := match[1]
	srcUrl := "https:" + src

	// Make the second GET request
	srcRes, err := http.Get(srcUrl)
	if err != nil {
		log.Fatalf("Error making GET request: %v\n", err)
	}
	defer srcRes.Body.Close()

	// Read the HTML content from the second response
	body, err = io.ReadAll(srcRes.Body)
	if err != nil {
		log.Fatalf("Error reading response body: %v\n", err)
	}

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
			<iframe src="` + src + `" allowfullscreen></iframe>
		</body>
		</html>`

	// Send the wrapped HTML back to the client
	c.Data(http.StatusOK, "text/html; charset=utf-8", []byte(wrappedHTML))
}
