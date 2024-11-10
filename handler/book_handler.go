package handler

import (
	"e-commerce-api/database"
	"e-commerce-api/models"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

//AddBook
func StoreBooks(c fiber.Ctx) error {
	book := new(models.Book)
	if err := c.Bind().Body(book); err != nil {
		return c.Status(400).JSON(err.Error())
	}

	database.DBConn.Create(&book)

	return c.Status(200).JSON(book)
}

func FindBooks(c fiber.Ctx) error {
	books := []models.Book{}

	database.DBConn.First(&books, c.Params("id"))

	return c.Status(200).JSON(books)
}

//AllBooks
func GetBooks(c fiber.Ctx) error {
	books := []models.Book{}

	database.DBConn.Find(&books)

	return c.Status(200).JSON(books)
}

//Update
func UpdateBooks(c fiber.Ctx) error {
	book := new(models.Book)
	if err := c.Bind().Body(book); err != nil {
		return c.Status(400).JSON(err.Error())
	}
	id, _ := strconv.Atoi(c.Params("id"))

	database.DBConn.Model(&models.Book{}).Where("id = ?", id).Update("title", book.Title)

	return c.Status(200).JSON("updated")
}

//Delete
func DeleteBooks(c fiber.Ctx) error {
	book := new(models.Book)

	id, _ := strconv.Atoi(c.Params("id"))

	database.DBConn.Where("id = ?", id).Delete(&book)

	return c.Status(200).JSON("deleted")
}
