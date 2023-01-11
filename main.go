package main

import (
	"github.com/izon-masuara/database"
	"github.com/izon-masuara/router"
)

func main() {
	r := router.SetUpRouter()
	router.Routers(r)
	database.Connect()
	r.Run(":3000")
}
