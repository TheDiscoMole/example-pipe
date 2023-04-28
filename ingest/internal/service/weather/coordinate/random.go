package coordinate

import (
    "math/rand"

    "github.com/TheDiscoMole/pipeline/service/ingest/pkg/model"
)

// sample coordinates randomly
func RandomBatch (seed int64, batchSize int) []model.Coordinate {
    coordinates := make([]model.Coordinate, batchSize)
    generator := rand.New(rand.NewSource(seed))

    for i := 0; i < batchSize; i++ {
        coordinates[i] = model.Coordinate{
            Latitude: generator.Float64() * 180 - 90,
            Longitude: generator.Float64() * 360 - 180,
        }
    }

    return coordinates
}
