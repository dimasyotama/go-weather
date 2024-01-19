package utils

import (
	"fmt"
	"io"
	"net/http"
	"encoding/json"
	"go-weather/config"
	"go-weather/dto"
	
)

func CurrentWeather(location string) (*dto_current_weather.WeatherData, error){
	if location == "" {
        return nil, fmt.Errorf("empty location provided")
    }

	api_key := config.Config("WEATHER_API_KEY") //Calling API KEY
	current_weather_url := fmt.Sprintf("https://api.weatherapi.com/v1/current.json?key=%s&q=%s", api_key, location) //CALLING ENDPOINT
	fmt.Println(current_weather_url)	
	response, err := http.Get(current_weather_url)
	
	if err != nil {
		return nil, fmt.Errorf("failed to make API call: %v", err)
	}

	defer response.Body.Close()

	// Read the response body
	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %v", err)
	}

	// Create an instance of the struct to hold the parsed JSON data
	var weatherData dto_current_weather.WeatherData

	// Unmarshal the JSON data into the struct
	if err := json.Unmarshal(body, &weatherData); err != nil {
		return nil, fmt.Errorf("error unmarshalling JSON: %v", err)
	}

	// Return the parsed data
	return &weatherData, nil
}
