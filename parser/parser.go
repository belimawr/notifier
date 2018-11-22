// Package parser is responsible for reading and parsing the messages
// currently it contains only a implementation to parse messages from
// git
package parser

import "github.com/belimawr/notifier/resources"

// Parser abstracts the reading and parsing of messages
// TODO: decide if reading of a message (e.g.: getting the output of
// `git log`) should be in a separeted interface/method
type Parser interface {
	// ParseMessage returns a parsed message ready to be sent on any
	// notification channel
	ParseMessage() (resources.Message, error)
}

// NewGit returns a Parser that reads and parses message from git logs
func NewGit() Parser {
	return git{}
}
