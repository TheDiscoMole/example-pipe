package openweathermap

import (
    "strconv"
    "fmt"

    "encoding/json"
    "net/http"

    "github.com/TheDiscoMole/pipeline/service/ingest/pkg/model"
)

func (c *Client) ForecastLocation (coordinate model.Coordinate) ([]*model.Forecast, error) {
    url := c.baseUrl + "/forecast" + c.urlArgs
    url  = fmt.Sprintf(url, c.appID, strconv.FormatFloat(coordinate.Latitude, 'f', -5, 64), strconv.FormatFloat(coordinate.Longitude, 'f', -5, 64), c.language, c.units)

    // get opm forecast
    response, err := c.client.Get(url)

    if err != nil {
        return nil, err
    }
    defer response.Body.Close()

    // validate status code
    if response.StatusCode != http.StatusOK {
        return nil, fmt.Errorf("OPM Status Code not OK: %d", response.StatusCode)
    }

    // decode forecast
    forecastResponse := opmResponse{}

    if err := json.NewDecoder(response.Body).Decode(&forecastResponse); err != nil {
        return nil, err
    }

    return forecastResponse.toForecasts(), nil
}

func (c *Client) ForecastLocationBatch (coordinates []model.Coordinate) ([]*model.Forecast, []error) {
    batch := make([]*model.Forecast, 0)
    errs := make([]error, len(coordinates))

    for i, coordinate := range(coordinates) {
        forecasts, err := c.ForecastLocation(coordinate)

        if err != nil {
            errs[i] = err
        } else {
            batch = append(batch, forecasts...)
        }
    }

    return batch, errs
}
