package post_deployment_functions

import (
	"github.com/DOC-SWAP/Docswap-backend/models"
	"gorm.io/gorm"
	"log"
	"slices"
)

func getParentIDByName(name string, dbCategories []models.Category) *uint {
	for _, category := range dbCategories {
		if category.Name == name {
			return &category.ID
		}
	}
	return nil
}

func CreateCategories(db *gorm.DB) {
	log.Print("Creating default Categories")
	// Insert categories
	categories := []models.Category{
		{Name: "Strata Corporations", Description: "Strata Plan\nBylaws\nAnnual Budget\nDepreciation Report\nMeeting Minutes\nAGM Special Assessments\nLitigation"},
		{Name: "Property Info", Description: "Architect Drawings\nFloorplan\nSurvey\nInspection Report\nRenovations\nCovenants / Easements\nValuations\nContracts"},
		{Name: "Mortgage Info", Description: "Appraisals\nRate Offers\nTerm"},
		{Name: "Rental Details", Description: "Suite rental\nCommercial Leases\nShort Term / Events\nVacation / Hotels"},
		{Name: "Content", Description: "Photo\nVideo\nWriting\nTemplates"},
		{Name: "Uncategorized", Description: "Uncategorized documents"},
	}

	var dbCategories []models.Category
	// get all the category names
	db.Select("name", "id", "parent_id").Where("parent_id IS NULL").Find(&dbCategories)

	var dbSubCategories []models.Category
	db.Select("name", "id", "parent_id").Where("parent_id IS NOT NULL").Find(&dbSubCategories)

	log.Println(dbCategories)
	log.Println(dbSubCategories)

	var categoryNames []string

	for _, category := range dbCategories {
		categoryNames = append(categoryNames, category.Name)
	}

	var subCategoryNames []string
	for _, category := range dbSubCategories {
		subCategoryNames = append(subCategoryNames, category.Name)
	}

	for _, category := range categories {
		if !slices.Contains(categoryNames, category.Name) {
			db.Create(&category)
		}
	}

	db.Select("name", "id").Find(&dbCategories)

	subCategories := []struct {
		Name        string
		Description string
		Parent      string
	}{
		// Content
		{"Photo", "Photos related to the content", "Content"},
		{"Video", "Videos related to the content", "Content"},
		{"Writing", "Written content documents", "Content"},
		{"Templates", "Templates for content", "Content"},

		// Rental Details
		{"Suite rental", "Details about suite rentals", "Rental Details"},
		{"Commercial Leases", "Details about commercial leases", "Rental Details"},
		{"Short Term / Events", "Details about short-term rentals and events", "Rental Details"},
		{"Vacation / Hotels", "Details about vacation rentals and hotels", "Rental Details"},

		// Property Info
		{"Architect Drawings", "Architectural drawings of the property", "Property Info"},
		{"Floorplan", "Floorplans of the property", "Property Info"},
		{"Survey", "Survey documents of the property", "Property Info"},
		{"Inspection Report", "Inspection reports of the property", "Property Info"},
		{"Renovations", "Renovation documents of the property", "Property Info"},
		{"Covenants / Easements", "Covenants and easements related to the property", "Property Info"},
		{"Appraisals / Valuations", "Appraisals and valuations of the property", "Property Info"},
	}
	var subCategoryObjects []models.Category

	for _, subCategory := range subCategories {
		if !slices.Contains(subCategoryNames, subCategory.Name) {
			parentId := getParentIDByName(subCategory.Parent, dbCategories)
			subCategoryObjects = append(subCategoryObjects, models.Category{
				Name:        subCategory.Name,
				Description: subCategory.Description,
				ParentID:    parentId,
			})
		}
	}

	db.Create(subCategoryObjects)

	log.Println("Default Categories created successfully.")

}
