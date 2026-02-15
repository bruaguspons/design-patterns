import { WeatherAPI, WeatherApp } from "./types";

export class WeatherAdapter implements WeatherApp {
    constructor(private weatherAPI: WeatherAPI) {}

    getTempF(): number {
        return this.weatherAPI.getTempC() * 9/5 + 32;
    }

    getHumidity(): number {
        return this.weatherAPI.getHumidity();
    }

    getWindSpeedMPH(): number {
        return this.weatherAPI.getWindSpeedKPH() * 0.621371;
    }
}