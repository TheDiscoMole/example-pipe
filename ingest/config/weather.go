package config

type Weather struct {
    OpenWeatherMap OpenWeatherMap
    OpenMeteo      OpenMeteo
    WeatherAPI     WeatherAPI

    SamplesPerDay  int
    BatchesPerHour int
}

type OpenWeatherMap struct {
    Key      string
    APILimit int
}

type OpenMeteo struct {
    APILimit int
}

type WeatherAPI struct {
    Key      string
    APILimit int
}

func loadWeather () Weather {
    return Weather{
        OpenWeatherMap: OpenWeatherMap{
            Key: getEnvAsString("WEATHER_OPENWEATHERMAP_KEY"),
            APILimit: getEnvAsInt("WEATHER_OPENWEATHERMAP_LIMIT"),
        },
        OpenMeteo: OpenMeteo{
            APILimit: getEnvAsInt("WEATHER_OPENMETEO_LIMIT"),
        },
        WeatherAPI: WeatherAPI{
            Key: getEnvAsString("WEATHER_WEATHERAPI_KEY"),
            APILimit: getEnvAsInt("WEATHER_WEATHERAPI_LIMIT"),
        },

        SamplesPerDay: getEnvAsInt("WEATHER_SAMPLES_PER_DAY"),
        BatchesPerHour: getEnvAsInt("WEATHER_BATCHES_PER_HOUR"),
    }
}
