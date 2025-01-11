package database

import (
	"api/models"
	"fmt"
	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var (
	DB *gorm.DB
)

// MysqlConnect initializes the connection to the MySQL database
func MysqlConnect() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Prepare the DSN (Data Source Name)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASSWORD"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME"))

	// Open the MySQL database connection using GORM
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // Avoid default transaction wrapping
		PrepareStmt:            true, // Enable statement preparation for better performance
	})
	if err != nil {
		log.Fatal("Failed to connect to database: ", err)
		os.Exit(2)
	}

	// Set up database connection pooling
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal("Failed to get db object: ", err)
	}
	// Set maximum idle and open connections, and maximum connection lifetime
	sqlDB.SetMaxIdleConns(10)
	sqlDB.SetMaxOpenConns(100)
	sqlDB.SetConnMaxLifetime(30 * time.Minute)

	// Perform AutoMigrate only for necessary models (or in development environment)
	log.Println("Connected to the database successfully")

	// AutoMigrate for models that need database schema changes
	err = db.AutoMigrate(
		&models.User{},
		&models.Product{},
		&models.Category{},
		&models.Order{},
		&models.OrderItem{},
		&models.Cart{},
		&models.UserInfo{},
		&models.Review{},
	)
	if err != nil {
		log.Fatal("AutoMigrate failed: ", err)
	}

	// Assign the DB object to the global variable
	DB = db
}
