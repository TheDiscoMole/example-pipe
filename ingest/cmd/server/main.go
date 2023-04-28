package main

import (
    "net/http"

    "github.com/TheDiscoMole/pipeline/service/ingest/config"
    "github.com/TheDiscoMole/pipeline/service/ingest/internal/server/router"

    "github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main () {
    configs := config.Load()

    server := chi.NewRouter()
    server.Use(middleware.Logger)
    server.Post("/weather", router.Weather(configs))

    http.ListenAndServe(":" + configs.Port, server)
}
