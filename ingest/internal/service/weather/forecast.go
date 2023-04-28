package weather

import (
    "context"
    "fmt"
    "time"

    "encoding/json"
    "github.com/TheDiscoMole/pipeline/service/ingest/internal/service/weather/coordinate"
)

func (c *Client) Forecast (ctx context.Context) error {
    currentTime := time.Now()
    seed := currentTime.Unix()

    // consume opm
    coordinates := coordinate.RandomBatch(seed, c.openweathermap.BatchSize)
    opmForecasts, opmErrs := c.openweathermap.ForecastLocationBatch(coordinates)

    // consume opm
    coordinates = coordinate.RandomBatch(seed, c.openmeteo.BatchSize)
    omForecasts, omErrs := c.openmeteo.ForecastLocationBatch(coordinates)

    // consume wapi
    coordinates = coordinate.RandomBatch(seed, c.weatherapi.BatchSize)
    wapiForecasts, wapiErrs := c.weatherapi.ForecastLocationBatch(coordinates)

    // merge results
    forecasts := append(opmForecasts, omForecasts...)
    forecasts  = append(forecasts, wapiForecasts...)

    errs := append(opmErrs, omErrs...)
    errs  = append(errs, wapiErrs...)

    // log errors
    for i, err := range errs {
        if err != nil {
            fmt.Printf("failed to fetch forecast at '%f.2' '%f.2': %s\n", coordinates[i].Latitude, coordinates[i].Longitude, err)
        }
    }

    // name batch
    filename := fmt.Sprintf("weather/%d", currentTime.Unix())
    data, _ := json.Marshal(forecasts)

    // save batch
    if err := c.Save(ctx, filename, data); err != nil {
        return err
    }

    // broadcast preprocessing task
    if err := c.Publish(ctx, "preprocess.weather", filename, nil); err != nil {
        return err
    }

    return nil
}
