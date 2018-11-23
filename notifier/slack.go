package notifier

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/belimawr/notifier/resources"
)

const (
	url    = "https://slack.com/api/chat.postMessage"
	msgFmt = "*%s* deployed `%s`: *%s*\n%s\n\n%s"
)

type slack struct {
	token       string
	channel     string
	application string
}

func (n slack) Notify(msg resources.Message) error {
	data := payload{
		Channel:  n.channel,
		Username: "Notifier Bot",
		IconURL:  "https://gophersource.com/img/mic-drop.png",
		Text: fmt.Sprintf(msgFmt,
			msg.Author,
			n.application,
			msg.Version,
			msg.Title,
			msg.Description,
		),
	}

	body, _ := json.Marshal(data)
	req, err := http.NewRequest("POST", url, bytes.NewReader(body))
	if err != nil {
		return err
	}

	req.Header.Add("Content-Type", "application/json; charset=utf-8")
	req.Header.Add("Authorization", "Bearer "+n.token)

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	respData := slackResp{}
	respBytes, _ := ioutil.ReadAll(resp.Body)
	if err := json.Unmarshal(respBytes, &respData); err != nil {
		return err
	}

	if !respData.OK {
		return fmt.Errorf("slack returned an error: %s", string(respBytes))
	}

	return nil
}

type payload struct {
	Channel  string `json:"channel"`
	Text     string `json:"text"`
	Username string `json:"username"`
	IconURL  string `json:"icon_url"`
}

type slackResp struct {
	OK bool `json:"ok"`
}
