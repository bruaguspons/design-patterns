package main

// Adapter
type WeatherAdapter struct {
	weatherAPI WeatherAPI
}

func NewWeatherAdapter(api WeatherAPI) *WeatherAdapter {
	return &WeatherAdapter{
		weatherAPI: api,
	}
}

func (w WeatherAdapter) GetTempF() float64 {
	return w.weatherAPI.GetTempC()*9/5 + 32
}

func (w WeatherAdapter) GetHumidity() float64 {
	return w.weatherAPI.GetHumidity()
}

func (w WeatherAdapter) GetWindSpeedMPH() float64 {
	return w.weatherAPI.GetWindSpeedKPH() * 0.621371
}
