package main

import "fmt"

func main() {
	api := ThirdPartyWeatherAPI{}
	weather := NewWeatherAdapter(api)

	fmt.Println("Temperature in Fahrenheit:", weather.GetTempF())
	fmt.Println("Humidity:", weather.GetHumidity())
	fmt.Println("Wind Speed in MPH:", weather.GetWindSpeedMPH())
}
