package config

type Storage struct {
    Bucket string
}

func loadStorage () Storage {
    return Storage{
        Bucket: getEnvAsString("STORAGE_BUCKET"),
    }
}
