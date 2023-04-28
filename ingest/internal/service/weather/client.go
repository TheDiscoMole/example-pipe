package weather

import (
    "github.com/TheDiscoMole/pipeline/service/ingest/config"
    "github.com/TheDiscoMole/pipeline/service/ingest/internal/service/weather/api/openweathermap"
    "github.com/TheDiscoMole/pipeline/service/ingest/internal/service/weather/api/openmeteo"
    "github.com/TheDiscoMole/pipeline/service/ingest/internal/service/weather/api/weatherapi"
    "github.com/TheDiscoMole/pipeline/service/ingest/pkg/publish"
    "github.com/TheDiscoMole/pipeline/service/ingest/pkg/repository"
)

type publishClient = publish.Client
type repositoryStorage = repository.Storage

type Client struct {
    *publishClient
    *repositoryStorage

    openweathermap *openweathermap.Client
    openmeteo      *openmeteo.Client
    weatherapi     *weatherapi.Client

    samplesPerDay   int
    batchesPerHour  int
}

func NewClient (configs *config.Config) *Client {
    return &Client{
        publishClient: publish.NewClient(configs),
        repositoryStorage: repository.NewStorage(configs),

        openweathermap: openweathermap.NewClient(configs),
        openmeteo: openmeteo.NewClient(configs),
        weatherapi: weatherapi.NewClient(configs),

        samplesPerDay: configs.Weather.SamplesPerDay,
        batchesPerHour: configs.Weather.BatchesPerHour,
    }
}
