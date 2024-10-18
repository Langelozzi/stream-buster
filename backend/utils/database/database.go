package database

import (
	"fmt"
	"github.com/STREAM-BUSTER/stream-buster/utils"
	_ "github.com/microsoft/go-mssqldb"
	"gorm.io/driver/sqlserver"
	"gorm.io/gorm"
	"log"
	"sync"
)

var (
	once     sync.Once
	instance *gorm.DB
)

func initialize() {
	// Build connection string
	connectionString := utils.GetEnvVariable("DB_CONNECTION_STRING")
	if connectionString == "" {
		log.Fatal("DB_CONNECTION_STRING is not set")
	}

	// create instance and Create connection pool
	// check to make sure the connection exists

	var err error
	instance, err = gorm.Open(sqlserver.Open(connectionString), &gorm.Config{})
	if err != nil {
		log.Fatal("Error creating connection pool: ", err.Error())
	}

	var result int
	if err := instance.Raw("SELECT 1").Scan(&result).Error; err != nil {
		log.Fatal("Error executing query: ", err.Error())
	}
	fmt.Printf("Database connection successful.\n")
}

// GetInstance returns the singleton instance of the database.
func GetInstance() *gorm.DB {
	once.Do(initialize)
	return instance
}
