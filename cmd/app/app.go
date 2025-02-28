package main

import (
	"asidikfauzi/xyz-multifinance-api/internal/database"
	"fmt"
)

func main() {
	fmt.Println("Starting server...")
	database.InitDatabase()
}
