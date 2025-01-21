package seeder

import (
	"api/models"
	"api/database"
	"log"
	"api/utils"
	"time"
)

func SeederSetup() {
	log.Println("Seeding database...")
	start.time.Now()
	//cek data seed
	models := []interface{}{&User{}, &Category, &Product{}}
	for _, model := range models {
		var count int64
		if err := db.Model(model).Count(&count).Error; err != nil {
			log.Println("Gagal menghitung data:", err)
			return err
		}
		if count > 0 {
			return nil // Skip jika data sudah ada
		}
	}
	
	// seeding user data
	password := utils.HashBycrypt("12345678")
	user := models.User {
		Name : "Yoga",
		Email : "yoga@gmail.com",
		Password: string(password),
	}
	
	if err := database.DB.Create(&user).Error; err != nil{
		log.Println(err)
	}
	// user info
	info := models.UserInfo{
		UserID: user.ID,
	}
	if err := database.DB.Create(&info).Error; err != nil{
		log.Println(err)
	}

	// seeding category data
	category := []models.Category{
		{Name : "Category 1"},
		{Name : "Category 2"},
	}
	

	for _,category := range category{
		if err := database.DB.Create(&category).Error; err != nil{
			log.Println(err)
		}
	}

	// seeding product data
	product := []models.Product{
		{Name : "Product 1", Description: "Cool Product", Price: 100000,Stock: 10, CategoryID: 1},
		{Name : "Product 2", Description: "Cool Product", Price: 200000,Stock: 20, CategoryID: 2},
	}

	for _,product := range product{
		if err := database.DB.Create(&product).Error; err != nil{
			log.Println(err)
		}
	}
	elapsed := time.Since(start)
	log.Println("Seeding database success..",elapsed.Milliseconds())


}
