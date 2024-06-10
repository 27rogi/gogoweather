package main

import (
	"fmt"
	"log"
	"os"

	"github.com/TwiN/go-color"
	_ "github.com/joho/godotenv/autoload"
)

var OWAPIKey = os.Getenv("OPENWEATHERAPI_KEY")

func main() {
	fmt.Println("\n🐹 GoGoWeather")
	if OWAPIKey == "" {
		log.Fatal("You did not provide OPENWEATHERAPI_KEY environment variable!")
	}

	loc := FetchLocation()
	fmt.Printf("\nShowing current weather for %v (%v)\n\n", color.InCyan(loc.City), color.OverCyan(loc.Country))

	weat := FetchWeather(loc.Latitude, loc.Longitude)
	fmt.Printf(
		"🌡️  Temperature: "+color.InYellow("%.1f")+" °C (feels like "+color.InPurple("%.0f")+" °C) | 💧 Humidity: "+color.InCyan("%.0f")+"%% \n💨 Wind Speed: %v m/s | ☁️  Cloud Percentage: %.1f%% \n🍃 Weather Condition is %v (%v)\n",
		weat.Stats.Temp, weat.Stats.TempFeel, weat.Stats.Humidity, weat.Wind.Speed, weat.Clouds.Percentage, color.InGreen(weat.Weather[0].Name), weat.Weather[0].Description,
	)
	fmt.Println("\nPress any ENTER to close the forecast...")
	fmt.Scanln()
}
