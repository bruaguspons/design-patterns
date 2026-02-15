class WeatherAdapter():
    def __init__(self, weather_api):
        self._weather_api = weather_api

    def get_temp_f(self) -> float:
        return self._weather_api.get_temp_c() * 9 / 5 + 32

    def get_humidity(self) -> float:
        return self._weather_api.get_humidity()

    def get_wind_speed_mph(self) -> float:
        return self._weather_api.get_wind_speed_kph() * 0.621371
