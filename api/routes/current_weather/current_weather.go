package current_weather_routes

import (
	"github.com/gofiber/fiber/v2"
	handler_current_weather "go-weather/api/handler/current_weather"
)
func SetupNoteRoutes(router fiber.Router){
	current_weather := router.Group("/weather")

	//Read one note
	current_weather.Get("/current-weather", handler_current_weather.GetCurrentWeather)


}