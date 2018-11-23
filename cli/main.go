package main

import (
	"os"

	"github.com/belimawr/notifier/notifier"
	"github.com/belimawr/notifier/parser"
)

func main() {
	git := parser.NewGit()

	m, err := git.ParseMessage()
	if err != nil {
		panic(err)
	}

	slack := notifier.NewSlack(
		os.Getenv("APP"),
		os.Getenv("CHANNEL"),
		os.Getenv("TOKEN"),
	)

	if err := slack.Notify(m); err != nil {
		panic(err)
	}
}
