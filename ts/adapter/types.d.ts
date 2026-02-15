export interface WeatherAPI {
    getTempC(): number;
    getHumidity(): number;
    getWindSpeedKPH(): number;
}

export interface WeatherApp {
    getTempF(): number;
    getHumidity(): number;
    getWindSpeedMPH(): number;
}