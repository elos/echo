package echo

import (
	"fmt"

	"github.com/elos/ehttp/serve"
	"github.com/subosito/twilio"
)

type Twilio interface {
	Send(to, body string) (*twilio.Message, *twilio.Response, error)
}

type Message struct {
	To, From, Body string
}

// Takes a http call from twilio and turns it into a message
// ExtractTwilioMessage
func Extract(c *serve.Conn) (*Message, error) {
	from := c.ParamVal("From")
	if from == "" {
		return nil, fmt.Errorf("Missing from parameter")
	}

	to := c.ParamVal("To")

	if to == "" {
		return nil, fmt.Errorf("Missing to parameter")
	}

	body := c.ParamVal("Body")

	if body == "" {
		return nil, fmt.Errorf("Missing body parameter")
	}

	return &Message{To: to, From: from, Body: body}, nil
}

func Handle(m *Message, t Twilio) {
	// just echo
	t.Send(m.From, m.Body)
}
