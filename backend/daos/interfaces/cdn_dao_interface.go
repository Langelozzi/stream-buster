package interfaces

type CDNDaoInterface interface {
	GetMovieContent(tmdbId string) (string, error)
	GetTVContent(tmdbId string, seasonNum int, episodeNum int) (string, error)
}
