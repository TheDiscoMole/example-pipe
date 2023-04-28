package openmeteo

import (
    "fmt"
    "time"

    "github.com/TheDiscoMole/pipeline/service/ingest/pkg/model"
)

type omResponse struct {
    Latitude                   float64  `json:"latitude"`
    Longitude                  float64  `json:"longitude"`
    Hourly                     omHourly `json:"hourly"`
}

type omHourly struct {
    Time                     []int      `json:"time"`
    RelativeHumidity         []float64  `json:"temperature_2m"`
    Temperature              []float64  `json:"relativehumidity_2m"`
    ApparentTemperature      []float64  `json:"apparent_temperature"`
    PrecipitationProbability []float64  `json:"precipitation_probability"`
    Rain                     []float64  `json:"rain"`
    Snowfall                 []float64  `json:"snowfall"`
    SurfacePressure          []float64  `json:"surface_pressure"`
    CloudCover               []float64  `json:"cloudcover"`
    Visibility               []float64  `json:"visibility"`
    WindSpeed                []float64  `json:"windspeed_10m"`
    WindDirection            []int      `json:"winddirection_10m"`
    WindGust                 []float64  `json:"windgusts_10m"`
}

func (f omResponse) toForecasts () []*model.Forecast {
    forecasts := make([]*model.Forecast, len(f.Hourly.Time))

    if len(forecasts) == 0 {
        fmt.Println("no openmeteo forecasts to map")
        return nil
    }

    timeForecasted := time.Now()

    for i, _ := range f.Hourly.Time {
        forecasts[i] = &model.Forecast{
            Source: "OpenMeteo",

            Time: f.Hourly.Time[i],
            TimeForecasted: int(timeForecasted.Unix()),

            Latitude: f.Latitude,
            Longitude: f.Longitude,
            City: "",
            Country: "",

            Temperature: f.Hourly.Temperature[i],
            TemperatureFeelsLike: f.Hourly.ApparentTemperature[i],

            Pressure: int(f.Hourly.SurfacePressure[i]),
            Humidity: float64(f.Hourly.RelativeHumidity[i]) / 100,
            PrecipitationProbability: f.Hourly.PrecipitationProbability[i],
            Cloudiness: float64(f.Hourly.CloudCover[i]) / 100,
            Visibility: float64(f.Hourly.Visibility[i]) / 24140,
            RainVolume: f.Hourly.Rain[i],
            SnowVolume: f.Hourly.Snowfall[i],
            WindSpeed: f.Hourly.WindSpeed[i],
            WindAngle: f.Hourly.WindDirection[i],
            WindGust: f.Hourly.WindGust[i],
        }
    }

    return forecasts
}
