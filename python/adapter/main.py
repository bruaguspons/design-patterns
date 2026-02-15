from WeatherAdapter import WeatherAdapter
from ThirdPartyWeatherAPI import ThirdPartyWeatherAPI

if __name__ == "__main__":
    weather = WeatherAdapter(ThirdPartyWeatherAPI())

    print("Temperature in Fahrenheit:", weather.get_temp_f())
    print("Humidity:", weather.get_humidity())
    print("Wind Speed in MPH:", weather.get_wind_speed_mph())
