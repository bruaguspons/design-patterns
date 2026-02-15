import { WeatherAPI } from "./types";

export class ThirdPartyWeatherAPI implements WeatherAPI {
    getTempC(): number {
        // Simulate fetching temperature in Celsius from a third-party API
        return 25; // Example temperature
    }

    getHumidity(): number {
        // Simulate fetching humidity from a third-party API
        return 60; // Example humidity
    }

    getWindSpeedKPH(): number {
        // Simulate fetching wind speed in KPH from a third-party API
        return 15; // Example wind speed
    }
}