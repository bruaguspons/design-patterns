import { ThirdPartyWeatherAPI } from "./ThirdPartyWeatherAPI";
import { WeatherAdapter } from "./WeatherAdapter";

const weather = new WeatherAdapter(new ThirdPartyWeatherAPI());

console.log("Temperature in Fahrenheit:", weather.getTempF());
console.log("Humidity:", weather.getHumidity());
console.log("Wind Speed in MPH:", weather.getWindSpeedMPH());