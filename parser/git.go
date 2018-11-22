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
	fields = []string{"%an", "%s", "%b", "%h"}
)

type git struct{}

func (g git) ParseMessage() (resources.Message, error) {
	format := fmt.Sprintf("--format=%s", strings.Join(fields, eol))

	rawGitLog, err := exec.Command("git", "log", "-n1", format).Output()
	if err != nil {
		return resources.Message{}, err
	}

	gitLog := string(rawGitLog)

	split := strings.Split(gitLog, eol)
	for i := 0; i < len(split); i++ {
		split[i] = strings.TrimSpace(split[i])
	}

	rawVersion, err := exec.Command("git", "describe", "--tags", split[3]).Output()
	if err != nil {
		return resources.Message{}, err
	}

	ver := strings.TrimSpace(string(rawVersion))

	msg := resources.Message{
		Author:      split[0],
		Title:       split[1],
		Description: split[2],
		Version:     ver,
	}

	return msg, nil
}
