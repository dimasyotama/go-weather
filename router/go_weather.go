package router

import (
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/logger"
	current_weather_routes "go-weather/api/routes/current_weather"

)

func SetupRoutes(app *fiber.App){
	api := app.Group("/api", logger.New())

	current_weather_routes.SetupNoteRoutes(api)

}