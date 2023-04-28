package publish

import (
    "github.com/TheDiscoMole/pipeline/service/ingest/config"
)

type Client struct {
    projectID string
}

func NewClient (configs *config.Config) *Client {
    return &Client{
        projectID: configs.ProjectID,
    }
}
