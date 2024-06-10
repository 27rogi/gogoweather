package main

import (
	"fmt"
	"log"

	"github.com/monaco-io/request"

)

const APIUrl = "https://api.openweathermap.org"

type Location struct {
	City        string  `json:"city"`
	Country     string  `json:"country_name"`
	CountryCode string  `json:"country_code"`
	Latitude    float32 `json:"latitude"`
	Longitude   float32 `json:"longitude"`
}

type Forecast struct {
	Weather []struct {
		Name        string `json:"main"`
		Description string `json:"description"`
	} `json:"weather"`
	Stats struct {
		Temp     float32 `json:"temp"`
		TempFeel     float32 `json:"feels_like"`
		Humidity float32 `json:"humidity"`
	} `json:"main"`
	Wind struct {
		Speed float32 `json:"speed"`
	} `json:"wind"`
	Clouds struct {
		Percentage float32 `json:"all"`
	} `json:"clouds"`
}

func FetchAPI[T any](path string) T {
	var userStruct T
	req := request.New().GET(APIUrl + path + "&appid=" + OWAPIKey).Send()
	if req.Code() != 200 {
		log.Fatal("There was an error fetching API data!")
	}
	req.ScanJSON(&userStruct)
	return userStruct
}

func FetchWeather(lat float32, lon float32) Forecast {
	forecast := FetchAPI[Forecast](fmt.Sprintf("/data/2.5/weather?lat=%.2f&lon=%.2f&units=metric", lat, lon))
	return forecast
}

func FetchLocation() Location {
	var loc Location
	req := request.New().GET("https://ipapi.co/json/").Send()
	if req.Code() != 200 {
		log.Fatal("There was an error fetching API data!")
	}
	req.ScanJSON(&loc)
	return loc
}
