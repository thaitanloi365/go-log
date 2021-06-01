package logger

import (
	"fmt"

	"github.com/parnurzeal/gorequest"
)

// SlackNotifier instance
type SlackNotifier struct {
	WebhookURL string
	ProxyURL   string
	Channel    string
}

func redirectPolicyFunc(req gorequest.Request, via []gorequest.Request) error {
	return fmt.Errorf("incorrect token (redirection)")
}

// NewSlackNotifier slack notifier
func NewSlackNotifier(webhookURL, channel string, proxyURL ...string) *SlackNotifier {
	var notifier = &SlackNotifier{
		WebhookURL: webhookURL,
		Channel:    channel,
	}

	if len(proxyURL) > 0 {
		notifier.ProxyURL = proxyURL[0]
	}

	return notifier
}

// Send send msg
func (slack *SlackNotifier) Send(title, body string) error {
	var payload = map[string]interface{}{
		"channel": slack.Channel,
		"blocks": []map[string]interface{}{
			{
				"type": "section",
				"text": map[string]interface{}{

					"type": "mrkdwn",
					"text": fmt.Sprintf("```%s\n%s```", title, body),
				},
			},
		},
	}
	request := gorequest.New().Proxy(slack.ProxyURL)
	resp, _, errs := request.
		Post(slack.WebhookURL).
		RedirectPolicy(redirectPolicyFunc).
		Send(payload).
		End()

	if len(errs) > 0 {
		return errs[0]
	}
	if resp.StatusCode >= 400 {
		return fmt.Errorf("Error sending msg. Status: %v", resp.Status)
	}

	return nil
}
