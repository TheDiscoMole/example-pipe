package openweathermap

import (
    "net/http"

    "github.com/TheDiscoMole/pipeline/service/ingest/config"
)

type Client struct {
    client   *http.Client

    baseUrl   string
    urlArgs   string
    appID     string
    language  string
    units     string

    APILimit  int
    BatchSize int
}

func NewClient (configs *config.Config) *Client {
    return &Client{
        client: http.DefaultClient,

        baseUrl: "https://api.openweathermap.org/data/2.5",
        urlArgs: "?appid=%s&lat=%s&lon=%s&lang=%s&units=%s",
        appID: configs.Weather.OpenWeatherMap.Key,
        language: "EN",
        units: "metric",

        APILimit: configs.Weather.OpenWeatherMap.APILimit,
        BatchSize: configs.Weather.OpenWeatherMap.APILimit / 31 / 24 / 60 * configs.Weather.BatchesPerHour,
    }
}
