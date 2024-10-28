package interfaces

import "github.com/STREAM-BUSTER/stream-buster/models/db"

type MediaDaoInterface interface {
	GetMediaById(mediaId uint) *db.Media
}
