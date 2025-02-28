package main

import (
	"asidikfauzi/xyz-multifinance-api/internal/database"
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"log"
)

func main() {
	db := database.InitDatabase()
	log.Println("Running AutoMigrate...")

	err := db.AutoMigrate(&model.Roles{}, &model.Users{}, &model.Consumers{}, &model.Limits{}, &model.Transactions{}, &model.Payments{})
	if err != nil {
		log.Fatal("Migration failed:", err)
	}

	log.Println("Migration completed successfully!")
}
