package interfaces

import (
	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/models/search"
)

type TagDaoInterface interface {
	GetAllTagsDao(includeDeleted bool, full bool) ([]models.Tag, error)
	GetTagDao(tagId int, includeDeleted bool, full bool) (*models.Tag, error)

	CreateTagDao(tag *models.Tag) (*models.Tag, error)
	CreateTagsBulkDao(tags []models.Tag) ([]models.Tag, error)

	SearchTagsDao(searchObj search.Search) ([]models.Tag, error)

	DeleteTagDao(tagId int) error
	DeleteTagsBulkDao(tags []models.Tag) error

	SoftDeleteTagDao(tagId int) error
	SoftDeleteTagsBulkDao(tags []models.Tag) error
}
