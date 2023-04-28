package weatherapi

import (
    "net/http"

    "github.com/TheDiscoMole/pipeline/service/ingest/config"
)

type Client struct {
    client  *http.Client

    baseUrl  string
    urlArgs  string
    key      string
    days     int
    language string

    APILimit  int
    BatchSize int
}

func NewClient (configs *config.Config) *Client {
    return &Client{
        client: http.DefaultClient,

        baseUrl: "http://api.weatherapi.com/v1",
        urlArgs: "?key=%s&q=%s,%s&days=%d&lang=%s",
        key: configs.Weather.WeatherAPI.Key,
        days: 3,
        language: "en",

        APILimit: configs.Weather.WeatherAPI.APILimit,
        BatchSize: configs.Weather.OpenWeatherMap.APILimit / 31 / 24 / 60 * configs.Weather.BatchesPerHour,
    }
}
