package services

import (
	iDao "github.com/STREAM-BUSTER/stream-buster/daos/interfaces"
	"github.com/STREAM-BUSTER/stream-buster/models/db"
)

type MediaService struct {
	dao iDao.MediaDaoInterface
}

func NewMediaService(dao iDao.MediaDaoInterface) *MediaService {
	return &MediaService{dao: dao}
}

func (service MediaService) CreateMedia(media *db.Media) (db.Media, error) {
	createdMedia := db.Media{
		ID:          1,
		TMDBID:      1100,
		Title:       "How I Met Your Mother",
		Overview:    "A father recounts to his children - through a series of flashbacks - the journey he and his four best friends took leading up to him meeting their mother.",
		PosterImage: "https://image.tmdb.org/t/p/w500/b34jPzmB0wZy7EjUZoleXOl2RRI.jpg",
		MediaTypeId: 0,
		MediaType: &db.MediaType{
			ID:          0,
			Name:        "tv",
			Description: "",
			DeletedAt:   nil,
			CreatedAt:   nil,
		},
		DeletedAt: nil,
		CreatedAt: nil,
		Genres:    nil,
	}
	return createdMedia, nil
}
