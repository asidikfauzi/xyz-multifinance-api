package main

import (
	"asidikfauzi/xyz-multifinance-api/internal/database"
	"asidikfauzi/xyz-multifinance-api/internal/model"
	"asidikfauzi/xyz-multifinance-api/internal/pkg/utils"
	"github.com/google/uuid"
	"log"
)

func main() {
	db := database.InitDatabase()
	// Seeding Roles
	roles := []model.Roles{
		{ID: uuid.New(), Name: "Admin"},
		{ID: uuid.New(), Name: "User"},
	}

	for i, role := range roles {
		db.FirstOrCreate(&roles[i], model.Roles{ID: role.ID})
	}

	// Seeding Users
	users := []model.Users{
		{ID: uuid.New(), Email: "admin@example.com", Password: utils.HashPassword("admin123"), RoleID: roles[0].ID},
		{ID: uuid.New(), Email: "user@example.com", Password: utils.HashPassword("user123"), RoleID: roles[1].ID},
	}

	for i, user := range users {
		db.FirstOrCreate(&users[i], model.Users{ID: user.ID})
	}

	log.Println("Seeding completed successfully!")
}
