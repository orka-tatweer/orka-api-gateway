package jobs

import (
	"bytes"
	"context"
	"encoding/json"
	"net/http"

	"github.com/charmbracelet/log"
	"github.com/redis/go-redis/v9"
)

type Event struct {
	SessionID  string `json:"session_id"`
	WebHookURL string `json:"webhook_url"`
	EventType  string `json:"event_type"`
}

// Monitor listens for events and broadcasts them through webhooks

func Listener(rdb *redis.Client, channel string) {
	sub := rdb.Subscribe(context.Background(), channel)

	subChan := sub.Channel()

	var payload Event

	for msg := range subChan {
		err := json.Unmarshal([]byte(msg.Payload), &payload)

		if err != nil {
			log.Warn(err)

		}

		data, err := json.Marshal(msg.Payload)

		if err != nil {
			log.Warn(err)
		}

		req, err := http.NewRequest("POST", payload.WebHookURL, bytes.NewBuffer(data))

		if err != nil {
			log.Warn(err)

		}
		// the request should be signed but didnt do that because of time
		req.Header.Set("Content-Type", "application/json")
		_, err = http.DefaultClient.Do(req)

		if err != nil {
			log.Warn(err)
		}
	}
}
