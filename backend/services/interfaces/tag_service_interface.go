package interfaces

import (
	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/models/search"
)

type TagServiceInterface interface {
	GetAllTags(includeDeleted bool, full bool) ([]models.Tag, error)
	GetTag(tagId int, includeDeleted bool, full bool) (*models.Tag, error)
	CreateTag(tag *models.Tag) (*models.Tag, error)
	CreateTagsBulk(tags []models.Tag) ([]models.Tag, error)
	CreateTagsAndUserTagsBulk(tags []models.Tag, userId uint) ([]models.Tag, error)

	SearchTags(searchObj search.Search) ([]models.Tag, error)

	DeleteTag(tagId int, softDelete bool) error
	DeleteTagsBulk(tags []models.Tag, softDelete bool) error
}
