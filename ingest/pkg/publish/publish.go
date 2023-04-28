package publish

import (
    "context"

    "cloud.google.com/go/pubsub"
)

func (c *Client) Publish (ctx context.Context, topic string, data string, attributes map[string]string) error {
    // create Google PubSub client
    client, err := pubsub.NewClient(ctx, c.projectID)

    if err != nil {
        return err
    }
    defer client.Close()

    // publish message
    publisher := client.Topic(topic)
    publishResult := publisher.Publish(
        ctx,
        &pubsub.Message{
            Data: []byte(data),
            Attributes: attributes,
        },
    )

    // handle failure
    if _, err := publishResult.Get(ctx); err != nil {
        return err
    }

    return nil
}
