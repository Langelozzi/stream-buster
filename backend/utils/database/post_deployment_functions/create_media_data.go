package post_deployment_functions

import (
	"errors"
	"fmt"
	"time"

	"github.com/STREAM-BUSTER/stream-buster/models"
	"github.com/STREAM-BUSTER/stream-buster/models/db"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CreateTestData creates test data for User, Media, MediaType, and CurrentlyWatching
func CreateTestData(database *gorm.DB) {
	// Step 1: Create test user if not exists (same as before)
	var existingUser models.User
	result := database.First(&existingUser, 1)
	if result.Error == nil {
		fmt.Println("User with ID 1 already exists. Skipping insert.")
	} else if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(result.Error)
	} else {
		createTestUser(database)
	}

	// Step 2: Create media types if they don't exist
	createMediaTypes(database)

	// Step 3: Create test media if not exists (same as before)
	var existingMedia db.Media
	result = database.First(&existingMedia, 1)
	if result.Error == nil {
		fmt.Println("Media with ID 1 already exists. Skipping insert.")
	} else if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		createTestMedia(database)
	}

	// Step 4: Create CurrentlyWatching entry if not exists (same as before)
	var existingCurrentlyWatching db.CurrentlyWatching
	result = database.Where("user_id = ? AND media_id = ?", 1, "1").First(&existingCurrentlyWatching)
	if result.Error == nil {
		fmt.Println("CurrentlyWatching entry for user 1 and media 1 already exists. Skipping insert.")
	} else if result.Error != nil && errors.Is(result.Error, gorm.ErrRecordNotFound) {
		createCurrentlyWatchingEntry(database)
	}
}

// createMediaTypes creates two rows in the MediaType table: TV and Movie
func createMediaTypes(database *gorm.DB) {
	// Check if "TV" media type exists
	var tvMediaType db.MediaType
	result := database.Where("name = ?", "TV").First(&tvMediaType)
	if result.Error == nil {
		fmt.Println("MediaType 'TV' already exists. Skipping insert.")
	} else if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(result.Error)
	} else {
		// Create "TV" media type
		tvMediaType := db.MediaType{
			Name:        "TV",
			Description: "Television series and shows",
			CreatedAt:   timePtr(time.Now()),
		}
		database.Create(&tvMediaType)
		fmt.Println("MediaType 'TV' created.")
	}

	// Check if "Movie" media type exists
	var movieMediaType db.MediaType
	result = database.Where("name = ?", "Movie").First(&movieMediaType)
	if result.Error == nil {
		fmt.Println("MediaType 'Movie' already exists. Skipping insert.")
	} else if result.Error != nil && !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		panic(result.Error)
	} else {
		// Create "Movie" media type
		movieMediaType := db.MediaType{
			Name:        "Movie",
			Description: "Feature films and movies",
			CreatedAt:   timePtr(time.Now()),
		}
		database.Create(&movieMediaType)
		fmt.Println("MediaType 'Movie' created.")
	}
}

// createTestUser creates a test user with ID 1
func createTestUser(db *gorm.DB) {
	unHashedPassword := "streambuster"
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(unHashedPassword), bcrypt.DefaultCost)
	if err != nil {
		fmt.Errorf("error hashing password")
	}

	user := models.User{
		ID:        1,
		Email:     "Admin@streambuster.com",
		FirstName: "Admin",
		LastName:  "LnameAdmin",
		Password:  string(hashedPassword),
		DeletedAt: nil,
	}
	db.Exec("SET IDENTITY_INSERT users ON")
	db.Create(&user)
	db.Exec("SET IDENTITY_INSERT users OFF")
	fmt.Println("User 'Admin' created.")
}

// createTestMedia creates a test media record with ID 1
func createTestMedia(database *gorm.DB) {
	media := db.Media{
		ID:          1,
		TMDBID:      1100,
		Title:       "How I Met You Mother",
		PosterImage: "https://image.tmdb.org/t/p/w500/pa4UM9lTaYLhi7RuBuPOejAoNfu.jpg",
		MediaTypeId: 1,
		CreatedAt:   timePtr(time.Now()),
	}
	database.Create(&media)

	media = db.Media{
		ID:          2,
		TMDBID:      767,
		Title:       "Harry Potter and the Half-Blood Prince",
		PosterImage: "https://image.tmdb.org/t/p/w500/pa4UM9lTaYLhi7RuBuPOejAoNfu.jpg",
		MediaTypeId: 2,
		CreatedAt:   timePtr(time.Now()),
	}
	database.Create(&media)
}

// createCurrentlyWatchingEntry creates a CurrentlyWatching record for User 1 and Media 1
func createCurrentlyWatchingEntry(database *gorm.DB) {
	currentlyWatching := db.CurrentlyWatching{
		UserID:        1,
		MediaId:       1,
		EpisodeNumber: 1,
		SeasonNumber:  1,
		CreatedAt:     timePtr(time.Now()),
	}
	database.Create(&currentlyWatching)

	currentlyWatching = db.CurrentlyWatching{
		UserID:    1,
		MediaId:   2,
		CreatedAt: timePtr(time.Now()),
	}
	database.Create(&currentlyWatching)
}

// timePtr is a helper function to create pointers to time.Time
func timePtr(t time.Time) *time.Time {
	return &t
}
