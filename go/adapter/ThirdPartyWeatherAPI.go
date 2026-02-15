package main

// Implementación concreta de la API externa
type ThirdPartyWeatherAPI struct{}

func (t ThirdPartyWeatherAPI) GetTempC() float64 {
	return 25
}

func (t ThirdPartyWeatherAPI) GetHumidity() float64 {
	return 60
}

func (t ThirdPartyWeatherAPI) GetWindSpeedKPH() float64 {
	return 15
}
