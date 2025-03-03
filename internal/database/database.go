package database

import (
	"api/internal/domain/models"
	"github.com/joho/godotenv"
	"gorm.io/gorm"
	"log"
	"os"
	"time"
)

var (
	DB *gorm.DB
)

// Setup initializes the connection to the MySQL database
func Setup() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Prepare the DSN (Data Source Name)
	user := os.Getenv("DB_USER")
	password := os.Getenv("DB_PASSWORD")
	host := os.Getenv("DB_HOST")
	port := os.Getenv("DB_PORT")
	name := os.Getenv("DB_NAME")
	driver := os.Getenv("DB_DRIVER")

	// Initialize db connection based on the driver
	var db *gorm.DB
	switch string(driver) {
	case "mysql":
		db = ConnectMySQL(user, password, host, port, name)
	case "pgsql":
		db = ConnectPostgres(user, password, host, port, name)
	default:
		log.Fatal("Unsupported DB_DRIVER. Please set it to either 'mysql' or 'postgres'")
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
	log.Println("Connected...")

	// AutoMigrate for models that need database schema changes
	err = db.AutoMigrate(
		&models.User{},
		&models.UserTemp{},
		&models.Product{},
		&models.Category{},
		&models.Order{},
		&models.OrderItem{},
		&models.Cart{},
		&models.UserInfo{},
		&models.Review{},
		&models.Permissions{},
		&models.Role{},
	)
	if err != nil {
		log.Fatal("AutoMigrate failed: ", err)
	}

	// Assign the DB object to the global variable
	DB = db
}
