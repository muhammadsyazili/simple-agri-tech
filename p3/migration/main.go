package main

import (
	"coding-test/p3/config"
	"coding-test/p3/models"
	"flag"
	"fmt"
)

func main() {
	fresh := flag.Bool("fresh", false, "Drop tables before migrating")
	flag.Parse()

	config.ConnectDatabase()
	db := config.DB

	if *fresh {
		fmt.Println("Dropping tables...")
		err := db.Migrator().DropTable(&models.User{}, &models.Spending{})
		if err != nil {
			panic("Failed to drop tables")
		}
		fmt.Println("Tables dropped successfully")
	}

	// Migrate the schema
	err := db.AutoMigrate(&models.User{}, &models.Spending{})
	if err != nil {
		panic("Failed to migrate database")
	}
	fmt.Println("Database migrated successfully")
}
