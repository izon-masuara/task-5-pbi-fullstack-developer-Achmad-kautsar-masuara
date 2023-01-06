package main

import "github.com/izon-masuara/router"

func main() {
	r := router.SetUpRouter()
	router.Routers(r)
	r.Run(":3000")
}
