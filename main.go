package main

import (
	"go-weather/utils"
	"github.com/gofiber/fiber/v2"
)

func main(){
	app := fiber.New()
	

	app.Listen(":2020")
}