package router

import (
    "fmt"

    "net/http"

    "github.com/TheDiscoMole/pipeline/service/ingest/config"
    "github.com/TheDiscoMole/pipeline/service/ingest/internal/service/weather"
)

func Weather (configs *config.Config) func (http.ResponseWriter, *http.Request) {
    return func (w http.ResponseWriter, r *http.Request) {
        ctx := r.Context()
        client := weather.NewClient(configs)

        if err := client.Forecast(ctx); err != nil {
            fmt.Println(err)
            w.WriteHeader(500)
        }
    }
}
