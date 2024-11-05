package interfaces

import "github.com/STREAM-BUSTER/stream-buster/models/db"

type MediaDaoInterface interface {
	CreateMedia(media *db.Media) error
	GetMediaById(id int64) (*db.Media, error)
}
