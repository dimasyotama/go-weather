package handler_current_weather

import (
	"log"
	"github.com/gofiber/fiber/v2"
	current_weather "go-weather/utils"
)

func GetCurrentWeather(c *fiber.Ctx) error {
	location :=c.Query("location")

	if location == "" {
        return c.Status(400).JSON(fiber.Map{"status": "error", "message": "Empty location provided", "data": nil})
    }

	weather_api, err := current_weather.CurrentWeather(location)

	log.Println(weather_api)
	if err != nil {
		return c.Status(404).JSON(fiber.Map{"status":"error", "message": "Data not found", "data": nil})
	}

	return c.JSON(fiber.Map{"status": "success", "message": "Notes Found", "data": weather_api})
}