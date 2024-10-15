package services

import (
	daosInterfaces "github.com/DOC-SWAP/Docswap-backend/daos/interfaces"
	"github.com/DOC-SWAP/Docswap-backend/models"
	"github.com/DOC-SWAP/Docswap-backend/models/search"
	"github.com/DOC-SWAP/Docswap-backend/services/interfaces"
)

type TagService struct {
	dao            daosInterfaces.TagDaoInterface
	userTagService interfaces.UserTagServiceInterface
}

func NewTagService(tagDao daosInterfaces.TagDaoInterface, userTagService interfaces.UserTagServiceInterface) *TagService {
	return &TagService{dao: tagDao, userTagService: userTagService}
}

func (service *TagService) GetAllTags(includeDeleted bool, full bool) ([]models.Tag, error) {
	return service.dao.GetAllTagsDao(includeDeleted, full)
}

func (service *TagService) GetTag(tagId int, includeDeleted bool, full bool) (*models.Tag, error) {
	return service.dao.GetTagDao(tagId, includeDeleted, full)
}

func (service *TagService) CreateTag(tag *models.Tag) (*models.Tag, error) {
	return service.dao.CreateTagDao(tag)
}

func (service *TagService) CreateTagsBulk(tags []models.Tag) ([]models.Tag, error) {
	return service.dao.CreateTagsBulkDao(tags)
}

func difference(a []models.Tag, b []models.Tag) []models.Tag {
	differenceResult := []models.Tag{}
	// for each element in a if its in b remove it from a
	// Convert b to a map for O(1) lookup
	tagBMap := make(map[string]bool)
	for _, tagB := range b {
		tagBMap[tagB.Name] = true
	}

	// Iterate over a and check if each tag is in the map
	for _, tagA := range a {
		if !tagBMap[tagA.Name] {
			differenceResult = append(differenceResult, tagA)
		}
	}
	return differenceResult
}

// takes an array of tags create the tags that need to be created. then create user tags for all the other tags.
// precondition: all the tags passed do not have a usertag associated with the user
func (service *TagService) CreateTagsAndUserTagsBulk(tags []models.Tag, userId uint) ([]models.Tag, error) {
	tagStringNames := []string{}
	for _, tag := range tags {
		tagStringNames = append(tagStringNames, tag.Name)
	}
	// use the search to check if the tags already exist
	searchObj := search.Search{
		Params: []search.Param{
			{
				Field:    "name",
				Operator: search.In,
				Value:    tagStringNames,
			},
		},
	}

	existingTags, err := service.dao.SearchTagsDao(searchObj)
	if err != nil {
		return nil, err
	}

	tagsToCreate := difference(tags, existingTags)
	var createdTags []models.Tag
	if len(tagsToCreate) != 0 {

		createdTags, err = service.dao.CreateTagsBulkDao(tagsToCreate)
		if err != nil {
			return nil, err
		}

	}

	// Create userTags for newTags and existingTags
	newUserTags := []models.UserTag{}
	for _, tag := range existingTags {
		newUserTags = append(newUserTags, models.UserTag{
			UserID: userId,
			TagID:  tag.ID,
		})
	}

	if len(newUserTags) == 0 {
		return createdTags, nil
	}

	_, err = service.userTagService.CreateUserTagsBulk(newUserTags)
	if err != nil {
		return nil, err
	}

	var tagsToReturn []models.Tag
	for _, tag := range createdTags {
		tagsToReturn = append(tagsToReturn, tag)
	}
	for _, tag := range existingTags {
		tagsToReturn = append(tagsToReturn, tag)
	}

	return tagsToReturn, nil
}

func (service *TagService) SearchTags(searchObj search.Search) ([]models.Tag, error) {
	return service.dao.SearchTagsDao(searchObj)
}

func (service *TagService) DeleteTag(tagId int, softDelete bool) error {
	if softDelete {
		return service.dao.SoftDeleteTagDao(tagId)
	}
	return service.dao.DeleteTagDao(tagId)
}

func (service *TagService) DeleteTagsBulk(tags []models.Tag, softDelete bool) error {
	if softDelete {
		return service.dao.SoftDeleteTagsBulkDao(tags)
	}
	return service.dao.DeleteTagsBulkDao(tags)
}
