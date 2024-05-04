package data

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

// Define the structure for your JSON data
type WeatherResponse struct {
	Properties WeatherProperties `json:"properties"`
}

type WeatherProperties struct {
	Periods []WeatherPeriod `json:"periods"`
}

type WeatherPeriod struct {
	Name             string `json:"name"`
	StartTime        string `json:"startTime"`
	EndTime          string `json:"endTime"`
	IsDaytime        bool   `json:"isDaytime"`
	Temperature      int    `json:"temperature"`
	TemperatureUnit  string `json:"temperatureUnit"`
	TemperatureTrend string `json:"temperatureTrend"`
	WindSpeed        string `json:"windSpeed"`
	WindDirection    string `json:"windDirection"`
	Icon             string `json:"icon"`
	ShortForecast    string `json:"shortForecast"`
	DetailedForecast string `json:"detailedForecast"`
}

type Temperature struct {
	Celsius    int
	Fahrenheit int
}

type WeatherForecast struct {
	Summary     string
	Temperature Temperature
}

func FahrenheitToCelsius(f int) int {
	return (f - 32) * 5 / 9
}

func GetWeatherForecast() WeatherForecast {
	// https://www.weather.gov/documentation/services-web-api#/default/gridpoint_forecast
	// https://api.weather.gov/gridpoints/OKX/33,35/forecast
	nycUri := "https://api.weather.gov/gridpoints/OKX/33,35/forecast"

	response, err := http.Get(nycUri)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return WeatherForecast{}
	}
	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		fmt.Printf("The HTTP request failed with error %s\n", err)
		return WeatherForecast{}
	}

	var data WeatherResponse
	err = json.Unmarshal(body, &data)
	if err != nil {
		fmt.Println("Error unmarshalling JSON: ", err)
		return WeatherForecast{}
	}

	forecast := WeatherForecast{
		Summary: data.Properties.Periods[0].DetailedForecast,
		Temperature: Temperature{
			Celsius:    FahrenheitToCelsius(data.Properties.Periods[0].Temperature),
			Fahrenheit: data.Properties.Periods[0].Temperature,
		},
	}

	return forecast
}
