package dwh

// https://discord.com/developers/docs/resources/webhook

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
)

type Webhook struct {
	Url    string
	Client http.Client
}

func NewWebhook(url string) Webhook {
	return Webhook{url, http.Client{}}
}

// sends a message to a webhook. Returns response status, error
func (webhook Webhook) Send(msg Message) (int, error) {
	// build msg json
	msgJSON, err := msg.toJSON()
	if err != nil {
		return 0, fmt.Errorf("failed to build json, %v", err)
	}

	// create request
	req, err := http.NewRequest(
		"POST",
		webhook.Url,
		bytes.NewBuffer(
			msgJSON,
		),
	)
	req.Header.Set("Content-Type", "application/json")
	if err != nil {
		return 0, fmt.Errorf("failed to create request, %v", err)
	}

	// send request
	resp, err := webhook.Client.Do(req)
	if err != nil {
		return 0, fmt.Errorf("failed to send request, %v", err)
	}

	// error checking
	if resp.StatusCode != 200 && resp.StatusCode != 204 {
		body, _ := io.ReadAll(resp.Body)
		log.Printf("their resp: %s", string(body))

		log.Printf("our req: %s", string(msgJSON))
	}

	// return status code
	return resp.StatusCode, nil
}

func (webhook Webhook) SendEmbed(emb Embed) (int, error) {
	msg := NewMessage("")
	msg.SetEmbed(emb)

	return webhook.Send(msg)
}
