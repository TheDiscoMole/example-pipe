package openweathermap

import (
    "fmt"
    "time"

    "github.com/TheDiscoMole/pipeline/service/ingest/pkg/model"
)

type opmResponse struct {
    Message     int            `json:"message"`
    List      []opmListElement `json:"list"`
    City        opmCity        `json:"city"`
}

type opmListElement struct {
    Dt         int             `json:"dt"`
    Main       opmMain         `json:"main"`
    Clouds     opmClouds       `json:"clouds"`
    Wind       opmWind         `json:"wind"`
    Visibility int             `json:"visibility"`
    Pop        float64         `json:"pop"`
    Rain       opmRain         `json:"rain"`
    Snow       opmSnow         `json:"snow"`
}

type opmCity struct {
    ID         int             `json:"id"`
    Name       string          `json:"name"`
    Coord      opmCoordinate   `json:"coord"`
    Country    string          `json:"country"`
    Population int             `json:"population"`
}

type opmCoordinate struct {
    Lat        float64        `json:"lat"`
    Lon        float64        `json:"lon"`
}

type opmMain struct {
    Temp       float64        `json:"temp"`
    FeelsLike  float64        `json:"feels_like"`
    Pressure   int            `json:"pressure"`
    SeaLevel   int            `json:"sea_level"`
    GrndLevel  int            `json:"grnd_level"`
    Humidity   int            `json:"humidity"`
}

type opmClouds struct {
    All        int            `json:"all"`
}

type opmWind struct {
    Speed      float64        `json:"speed"`
    Deg        int            `json:"deg"`
    Gust       float64        `json:"gust"`
}

type opmRain struct {
    ThreeH     float64        `json:"3h"`
}

type opmSnow struct {
    ThreeH     float64        `json:"3h"`
}

func (f opmResponse) toForecasts () []*model.Forecast {
    forecasts := make([]*model.Forecast, len(f.List))

    if len(forecasts) == 0 {
        fmt.Println("no openweathermap forecasts to map")
        return nil
    }

    timeForecasted := time.Now()

    for i, listElement := range f.List {
        forecasts[i] = &model.Forecast{
            Source: "OpenWeatherMap",

            Time: listElement.Dt,
            TimeForecasted: int(timeForecasted.Unix()),

            Latitude: f.City.Coord.Lat,
            Longitude: f.City.Coord.Lon,
            City: f.City.Name,
            Country: f.City.Country,

            Temperature: listElement.Main.Temp,
            TemperatureFeelsLike: listElement.Main.FeelsLike,

            Pressure: listElement.Main.Pressure,
            Humidity: float64(listElement.Main.Humidity) / 100,
            PrecipitationProbability: listElement.Pop,
            Cloudiness: float64(listElement.Clouds.All) / 100,
            Visibility: float64(listElement.Visibility) / 10000,
            RainVolume: listElement.Rain.ThreeH / 3,
            SnowVolume: listElement.Snow.ThreeH / 3,
            WindSpeed: listElement.Wind.Speed,
            WindAngle: listElement.Wind.Deg,
            WindGust: listElement.Wind.Gust,
        }
    }

    return forecasts
}
