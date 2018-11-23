package notifier

import "github.com/belimawr/notifier/resources"

// Notifier abstracts the notification of a message
type Notifier interface {
	// Notify receives a message and sends it to one notification channel
	Notify(msg resources.Message) error
}

// NewSlack returns a implementation fo Notifier that sends messages
// on Slack using Slack API `chat.postMessage`
func NewSlack(app, channel, token string) Notifier {
	return slack{
		application: app,
		channel:     channel,
		token:       token,
	}
}
