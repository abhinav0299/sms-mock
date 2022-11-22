package twilio

import (
	"errors"

	"github.com/subosito/twilio"
)

type Twilio struct {
	Config map[string]string
}

func (t Twilio) Send(to, message string) error {
	var (
		accountSid, authToken, from string
	)

	if t.Config["accountSid"] != "" {
		accountSid = t.Config["accountSid"]
	}
	if t.Config["authToken"] != "" {
		authToken = t.Config["authToken"]
	}
	if t.Config["from"] != "" {
		from = t.Config["from"]
	}
	if accountSid == "" || authToken == "" || from == "" {
		return errors.New("invalid creds")
	}
	c := twilio.NewClient(accountSid, authToken, nil)
	params := twilio.MessageParams{
		Body: message,
	}
	_, _, err := c.Messages.Send(from, to, params)
	if err != nil {
		return errors.New("couldn't send sms")
	}
	return nil
}
