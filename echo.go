package echo

import (
	"fmt"
	"strings"

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

func parse(s string) (cmd string, bdy string) {
	firstSpace := strings.Index(s, " ")

	if firstSpace == -1 {
		return
	}

	cmd = strings.ToLower(s[0:firstSpace])
	bdy = s[firstSpace+1 : len(s)-1]
	return
}

func Handle(m *Message, t Twilio) {
	cmd, body := parse(m.Body)
	switch cmd {
	case "start":
		start(m, body, t)
		return
	default:
		break
	}
	// just echo
	t.Send(m.From, m.Body)
}

func start(m *Message, parseBody string, t Twilio) {
	t.Send(m.From, fmt.Sprintf("Starting..."))
	t.Send(m.From, fmt.Sprintf("Started %s", parseBody))
}
