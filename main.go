package main

import (
	"github.com/izon-masuara/database"
	"github.com/izon-masuara/router"
	"github.com/joho/godotenv"
)

func main() {
	if err := godotenv.Load(".env"); err != nil {
		panic("Error .env file required")
	}
	r := router.SetUpRouter()
	router.Routers(r)
	database.Connect()
	r.Run(":3000")
}
