package post_deployment_functions

import (
	"gorm.io/gorm"
	"log"
)

func CreateUserTotalRequestCountView(db *gorm.DB) error {
	// Check if the view already exists
	var count int
	checkQuery := `
    SELECT COUNT(*)
    FROM information_schema.views
    WHERE table_schema = 'public' AND table_name = 'user_total_request_count';
    `

	err := db.Raw(checkQuery).Scan(&count).Error
	if err != nil {
		return err
	}

	if count > 0 {
		log.Println("View 'user_total_request_count' already exists.")
		return nil
	}

	// SQL query to create the view
	createQuery := `
	CREATE VIEW user_total_request_counts AS
	SELECT 
		user_id,
		SUM(request_count) AS total_request_count
	FROM 
		user_endpoint_usages
	GROUP BY 
		user_id;
    `

	// Execute the query to create the view
	if err := db.Exec(createQuery).Error; err != nil {
		return err
	}

	log.Println("View 'user_total_request_count' created successfully.")
	return nil
}
