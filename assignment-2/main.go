package main

import (
	"assignment-2_septyan/config"
	"assignment-2_septyan/controllers"
	"assignment-2_septyan/routers"
)

var PORT = ":8081"

func main() {
	db := config.StartDB()
	controller := controllers.New(db)

	routers.StartServer(controller).Run(PORT)
}
