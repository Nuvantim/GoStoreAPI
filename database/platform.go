package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// mysql platform
func ConnectMySQL(user, password, host, port, name string) *gorm.DB {
	dsn := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + name + "?charset=utf8mb4&parseTime=True&loc=Local"
	// Open Connection Mysql
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // Avoid default transaction wrapping
		PrepareStmt:            true, // Enable statement preparation for better performance
	})
	if err != nil {
		log.Fatal("Failed to connect to MySQL database: ", err)
		os.Exit(2)
	}
	return db
}

func ConnectPostgres(user, password, host, port, name string) *gorm.DB {
	dsn := "host=" + host + " user=" + user + " password=" + password + " dbname=" + name + " port=" + port + " sslmode=disable TimeZone=Asia/Jakarta"
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true,
		PrepareStmt:            true,
	})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL database: ", err)
		os.Exit(2)
	}
	return db
}
