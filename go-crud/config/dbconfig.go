package config

import (
	"fmt"
	"go/go-crud/modal"
	"os"
	"time"

	"log"

	"github.com/cenkalti/backoff/v4"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConfigureDatabase() {

	dbUser := os.Getenv("DB_USER")
	dbPassword := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// GORM configuration
	var db *gorm.DB
	var err error

	// refer https://github.com/go-sql-driver/mysql#dsn-data-source-name for details
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		dbUser, dbPassword, dbHost, dbPort, dbName)

	// Define backoff parameters
	bo := backoff.NewExponentialBackOff()
	bo.InitialInterval = 100 * time.Millisecond
	bo.MaxInterval = 10 * time.Second
	bo.MaxElapsedTime = 5 * time.Minute // Maximum total time for retries

	// Retry function for establishing database connection
	operation := func() error {
		db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
		if err != nil {
			log.Printf("Failed to connect to database: %v\n", err)
			return err
		}

		// Check if the connection is alive
		sqlDB, err := db.DB()
		if err != nil {
			log.Printf("Failed to get underlying sql.DB: %v\n", err)
			db = nil
			return err
		}

		err = sqlDB.Ping()
		if err != nil {
			log.Printf("Failed to ping database: %v\n", err)
			db = nil
			return err
		}

		// Connection successful
		log.Printf("Database connection established successfully.")
		return nil
	}

	// Execute the operation with backoff retries
	err = backoff.Retry(operation, bo)
	if err != nil {
		log.Panicf("Failed to establish database connection: %v", err)
	}

	// Perform any migrations or setup here
	db.AutoMigrate(&modal.User{}) // Example: Auto migrate User model

	// Assign the connected database to the global DB variable
	DB = db
}
