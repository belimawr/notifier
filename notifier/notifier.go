package notifier

import "github.com/belimawr/notifier/resources"

// Notifier abstracts the notification of a message
type Notifier interface {
	// Notify receives a message and sends it to one notification channel
	Notify(msg resources.Message) error
}
