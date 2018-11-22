package parser

import (
	"fmt"
	"os/exec"
	"strings"

	"github.com/belimawr/notifier/resources"
)

var (
	// eol is the end-of-line string used to separate different fields on the
	// git log output
	eol = "<<EOL>>\n"

	// fields is a list of fields to be extracted from git log
	fields = []string{"%an", "%s", "%b"}
)

type git struct{}

func (g git) ParseMessage() (resources.Message, error) {
	format := fmt.Sprintf("--format=%s", strings.Join(fields, eol))

	rawOutput, err := exec.Command("git", "log", "-n1", format).Output()
	if err != nil {
		return resources.Message{}, err
	}

	output := string(rawOutput)

	split := strings.Split(output, eol)
	for i := 0; i < len(split); i++ {
		split[i] = strings.TrimSpace(split[i])
	}

	msg := resources.Message{
		Author:      split[0],
		Title:       split[1],
		Description: split[2],
	}

	return msg, nil
}
