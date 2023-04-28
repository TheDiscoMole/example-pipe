package repository

import (
    "context"

    "github.com/TheDiscoMole/pipeline/service/ingest/config"

    "cloud.google.com/go/storage"
)

type Storage struct {
    ProjectID string
    Bucket    string
}

func NewStorage (configs *config.Config) *Storage {
    return &Storage{
        ProjectID: configs.ProjectID,
        Bucket: configs.Storage.Bucket,
    }
}

func (s *Storage) Save (ctx context.Context, filename string, data []byte) error {
    client, err := storage.NewClient(ctx)

    if err != nil {
        return err
    }
    defer client.Close()

    bucket := client.Bucket(s.Bucket)
    object := bucket.Object("ingest/" + filename)
    writer := object.NewWriter(ctx)

    if _, err := writer.Write(data); err != nil {
        return err
    }
    if err := writer.Close(); err != nil {
        return err
    }

    return nil
}
