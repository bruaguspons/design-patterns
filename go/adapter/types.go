package main

// Interfaz que representa la API externa
type WeatherAPI interface {
	GetTempC() float64
	GetHumidity() float64
	GetWindSpeedKPH() float64
}

// Interfaz que espera nuestra aplicación
type WeatherApp interface {
	GetTempF() float64
	GetHumidity() float64
	GetWindSpeedMPH() float64
}
