package interfaces

import "github.com/STREAM-BUSTER/stream-buster/models/db"

type MediaDaoInterface interface {
	CreateMedia(media *db.Media) (*db.Media, error)
	GetMediaById(id int64) (*db.Media, error)
	GetMediaByTMDBId(id int64) (*db.Media, error)
}
