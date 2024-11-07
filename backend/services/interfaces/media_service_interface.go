package interfaces

import (
	"github.com/STREAM-BUSTER/stream-buster/models/db"
)

type MediaServiceInterface interface {
	CreateMedia(media *db.Media) error
	GetMediaById(id int64) (*db.Media, error)
}