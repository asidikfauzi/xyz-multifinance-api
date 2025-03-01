package main

import (
	"asidikfauzi/xyz-multifinance-api/internal/config"
	"asidikfauzi/xyz-multifinance-api/internal/database"
	"asidikfauzi/xyz-multifinance-api/internal/server"
	"fmt"
)

func main() {
	fmt.Println("Starting server...")
	database.InitDatabase()

	s := server.InitializedServer()

	err := s.Engine.Run(fmt.Sprintf(":%s", config.Env("APP_PORT")))
	if err != nil {
		return
	}
}
