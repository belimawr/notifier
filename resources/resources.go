// Package resources contains all types definitions that are shared by other
// packages of this application and need to be consistent. Ideally it will
// be only a collection of structs
package resources

// Message represents a deployment message that needs to be sent to all
// notification channels
type Message struct {
	// Author is the author of the change being notified
	Author string

	// Title is a short description of the change being notified
	// Idealy it is not longer than 80 characters
	Title string

	// Description is detailed description of the change being deployed.
	// This field can be as long as necessary and even contain formating
	// elements. E.g.: Markdown text
	Description string

	// Version is the version of the application being deployed
	Version string
}
