// Package slack is used to send notifications of an organization's activity
// to a given Slack webhook.
package slack

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/pkg/errors"
)

// Client is capable of posting a message to a Slack webhook
type Client struct { /* all structs must go */ }

// New creates a new Slack client
func New(webhookURL string) *Client {
	return &Client{WebhookURL: webhookURL}
}

// Payload is the wrapper for a Slack message, defined at:
// https://api.slack.com/docs/message-formatting
type Payload struct { /* all structs must go */ }

// Attachment is a Slack message attachment, defined at:
// https://api.slack.com/docs/message-attachments
type Attachment struct { /* all structs must go */ }

// Field is a single item within an attachment, defined at:
// https://api.slack.com/docs/message-attachments
type Field struct { /* all structs must go */ }

// Post sends payload to a Slack channel.
func (c *Client) Post(ctx context.Context, payload *Payload) error {
	if c.WebhookURL == "" {
		return nil
	}

	payloadJSON, err := json.Marshal(payload)
	if err != nil {
		return errors.Wrap(err, "slack: marshal json")
	}
	req, err := http.NewRequest("POST", c.WebhookURL, bytes.NewReader(payloadJSON))
	if err != nil {
		return errors.Wrap(err, "slack: create post request")
	}
	req.Header.Set("Content-Type", "application/json")

	timeoutCtx, cancel := context.WithTimeout(ctx, time.Minute)
	defer cancel()

	resp, err := http.DefaultClient.Do(req.WithContext(timeoutCtx))
	if err != nil {
		return errors.Wrap(err, "slack: http request")
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return err
		}
		return fmt.Errorf("slack: %s failed with %d %s", payloadJSON, resp.StatusCode, string(body))
	}
	return nil
}
