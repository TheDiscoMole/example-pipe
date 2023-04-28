package config

import (
    "log"
    "strconv"
    "os"

    "github.com/joho/godotenv"
)

type Config struct {
    Port      string
    Mode      string
    ProjectID string

    Storage   Storage
    Weather   Weather
}

func Load () *Config {
    // load environment file
    err := godotenv.Load(".env")

    if err != nil {
        log.Fatal("failed to load .env file")
    }

    // build configs
    configs := Config{
        Port: getEnvAsString("PORT"),
        Mode: getEnvAsString("MODE"),
        ProjectID: getEnvAsString("PROJECT_ID"),

        Storage: loadStorage(),
        Weather: loadWeather(),
    }

    return &configs
}

// load environment variable as string
func getEnvAsString (key string) string {
    value := os.Getenv(key)

    if value == "" {
        log.Fatalf("failed to load '%s' environment variable as string", key)
    }

    return value
}

// load environment variable as int
func getEnvAsInt (key string) int {
    value, err := strconv.Atoi(getEnvAsString(key))

    if err != nil {
        log.Fatalf("failed to load '%s' environment variable as int", key)
    }

    return value
}
