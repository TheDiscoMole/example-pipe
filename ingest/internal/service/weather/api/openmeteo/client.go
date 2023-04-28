package openmeteo

import (
    "net/http"

    "github.com/TheDiscoMole/pipeline/service/ingest/config"
)

type Client struct {
    client       *http.Client

    baseUrl       string
    urlArgs       string
    forecastDays  int
    windSpeedUnit string
    timeFormat    string
    measurements  string

    APILimit      int
    BatchSize     int
}

func NewClient (configs *config.Config) *Client {
    return &Client{
        client: http.DefaultClient,

        baseUrl: "https://api.open-meteo.com/v1",
        urlArgs: "?latitude=%s&longitude=%s&forecast_days=%d&windspeed_unit=%s&timeformat=%s&hourly=%s",
        forecastDays: 7,
        windSpeedUnit: "ms",
        timeFormat: "unixtime",
        measurements: "temperature_2m,relativehumidity_2m,apparent_temperature,precipitation_probability,rain,snowfall,surface_pressure,cloudcover,visibility,windspeed_10m,winddirection_10m,windgusts_10m",

        APILimit: configs.Weather.OpenMeteo.APILimit,
        BatchSize: configs.Weather.OpenMeteo.APILimit / 31 / 24 / 60 * configs.Weather.BatchesPerHour,
    }
}
