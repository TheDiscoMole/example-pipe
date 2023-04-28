package weatherapi

import (
    "fmt"
    "math"
    "time"

    "github.com/TheDiscoMole/pipeline/service/ingest/pkg/model"
)

type wapiResponse struct {
    Location      wapiLocation    `json:"location"`
    Forecast      wapiForecast    `json:"forecast"`
    Current       wapiHour        `json:"current"`
}

type wapiLocation struct {
    Name          string          `json:"name"`
    Country       string          `json:"country"`
    Lat           float64         `json:"lat"`
    Lon           float64         `json:"lon"`
}

type wapiForecast struct {
    ForecastDay []wapiForecastDay `json:"forecastday"`
}

type wapiForecastDay struct {
    DateEpoch     int             `json:"date_epoch"`
    Hour        []wapiHour        `json:"hour"`
}

type wapiHour struct {
    TimeEpoch     int             `json:"time_epoch"`
    TempC         float64         `json:"temp_c"`
    WindKph       float64         `json:"wind_kph"`
    GustKph       float64         `json:"gust_kph"`
    WindDegree    int             `json:"wind_degree"`
    PressureMb    float64         `json:"pressure_mb"`
    PrecipMm      float64         `json:"precip_mm"`
    Humidity      float64         `json:"humidity"`
    Cloud         float64         `json:"cloud"`
    FeelslikeC    float64         `json:"feelslike_c"`
    WillItRain    float64         `json:"will_it_rain"`
    ChanceOfRain  float64         `json:"chance_of_rain"`
    WillItSnow    float64         `json:"will_it_snow"`
    ChanceOfSnow  float64         `json:"chance_of_snow"`
    VisKm         float64         `json:"vis_km"`
}

func (f wapiResponse) toForecasts () []*model.Forecast {
    if len(f.Forecast.ForecastDay) == 0 {
        fmt.Println("no weatherapi forecasts to map")
        return nil
    }

    forecasts := make([]*model.Forecast, 0)
    timeForecasted := time.Now()

    for _, day := range f.Forecast.ForecastDay {
        for _, hour := range day.Hour {
            forecasts = append(forecasts, &model.Forecast{
                Source: "WeatherAPI",
                
                Time: hour.TimeEpoch,
                TimeForecasted: int(timeForecasted.Unix()),

                Latitude: f.Location.Lat,
                Longitude: f.Location.Lon,
                City: f.Location.Name,
                Country: f.Location.Country,

                Temperature: hour.TempC,
                TemperatureFeelsLike: hour.FeelslikeC,

                Pressure: int(hour.PressureMb),
                Humidity: hour.Humidity / 100,
                PrecipitationProbability: math.Max(float64(hour.ChanceOfRain) / 100, float64(hour.ChanceOfSnow) / 100),
                Cloudiness: hour.Cloud / 100,
                Visibility: hour.VisKm / 10,
                RainVolume: hour.PrecipMm * hour.WillItRain,
                SnowVolume: hour.PrecipMm * hour.WillItSnow,
                WindSpeed: hour.WindKph / 3.6,
                WindAngle: hour.WindDegree,
                WindGust: hour.GustKph / 3.6,
            })
        }
    }

    return forecasts
}
