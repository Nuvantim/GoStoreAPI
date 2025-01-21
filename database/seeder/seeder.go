package seeder

import (
	"log"
)

func SeederSetup() {
	log.Println("Seeding database...")

	// seeding user data
	seed_User()

	// seeding category data
	seed_Category()

	// seeding product data
	seed_Product()

	log.Println("Seeding database success..")

}
