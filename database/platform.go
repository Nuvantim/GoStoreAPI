package database

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"os"
)

// mysql platform
func ConnectMySQL(user, password, host, port, name string) *gorm.DB {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		user,
		password,
		host,
		port,
		name)
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
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		host,
		user,
		password,
		name,
		port)
	// Open Connection Postgres
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		SkipDefaultTransaction: true, // Avoid default transaction wrapping
		PrepareStmt:            true, // Enable statement preparation for better performance
	})
	if err != nil {
		log.Fatal("Failed to connect to PostgreSQL database: ", err)
		os.Exit(2)
	}
	return db
}
