package model

type Forecast struct {
    Source                   string  `json:"source"`

    Time                     int     `json:"time"`
    TimeForecasted           int     `json:"time_forecasted"`

    Latitude                 float64 `json:"latitude"`
    Longitude                float64 `json:"longitude"`
    City                     string  `json:"city"`
    Country                  string  `json:"country"`

    Temperature              float64 `json:"temperature"`
    TemperatureFeelsLike     float64 `json:"temperature_feels_like"`

    Pressure                 int     `json:"pressure"`
    Humidity                 float64 `json:"humidity"`
    PrecipitationProbability float64 `json:"precipitation_probability"`
    Cloudiness               float64 `json:"cloudiness"`
    Visibility               float64 `json:"visibility"`
    RainVolume               float64 `json:"rain_volume"`
    SnowVolume               float64 `json:"snow_volume"`
    WindSpeed                float64 `json:"wind_speed"`
    WindAngle                int     `json:"wind_angle"`
    WindGust                 float64 `json:"wind_gust"`
}
