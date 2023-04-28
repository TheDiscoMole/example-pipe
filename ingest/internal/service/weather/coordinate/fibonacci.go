package coordinate

import (
    "math"
    "time"

    "math/rand"

    "github.com/TheDiscoMole/pipeline/service/ingest/pkg/model"
)

// not sure if this is even nessecary

// sample coordinates evenly accross globe
// https://stackoverflow.com/questions/9600801/evenly-distributing-n-points-on-a-sphere
func FibonacciBatch (seed int64, samples int, offset int, batchSize int) []model.Coordinate {
    coordinates := make([]model.Coordinate, batchSize)
    generator := rand.New(rand.NewSource(seed))

    phi := math.Pi * (math.Sqrt(5) - 1)

    dLatitude := generator.Float64() * 180 - 90
    dLongitude := generator.Float64() * 360 - 180

    for i := 0; i < batchSize; i++ {
        sample := float64(i + offset)
        theta := phi * sample

        y := 1 - (sample / float64(samples - 1)) * 2
        x := math.Cos(theta) * math.Sqrt(1 - y * y)

        latitude := math.Asin(math.Sin(theta)) * 180 / math.Pi
        longitude := math.Atan2(y, x) * 180 / math.Pi

        coordinates[i] = model.Coordinate{
            Latitude: math.Mod(latitude + dLatitude, 90),
            Longitude: math.Mod(longitude + dLongitude, 180),
        }
    }

    return coordinates
}

func FibonacciArgumentHelper (currentTime time.Time, apiLimit int, samplesPerDay int, batchesPerHour int) (int, int, int) {
    batchSize := apiLimit / 31 / 24 / batchesPerHour
    samples := batchSize * batchesPerHour * 24 / samplesPerDay

    hours, minutes, _ := currentTime.Clock()
    offset := (hours % (24 / samplesPerDay)) * batchesPerHour
    offset += minutes / (60 / batchesPerHour)
    offset *= batchSize

    return samples, offset, batchSize
}
