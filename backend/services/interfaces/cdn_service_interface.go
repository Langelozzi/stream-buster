package interfaces

type CDNServiceInterface interface {
	GetMovieContent(tmdbId string) (string, error)
	GetTVContent(tmdbId string, seasonNum int, episodeNum int) (string, error)
	CheckContentExists(tmdbId string, isTV bool) bool
}
