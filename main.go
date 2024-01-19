package main

import (
	"github.com/gofiber/fiber/v2"
	"go-weather/router"
	"go-weather/config"
)

func main(){
	app := fiber.New()

	//pointing router to router group
	router.SetupRoutes(app)

	//assign port from .env
	port := config.Config("PORT")
	app.Listen(":" + port)
}