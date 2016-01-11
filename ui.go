package echo

import (
	"fmt"
	"time"
)

// A TextUI is used for making command line interfaces
// more suitable for a medium in which you can only communicate
// strings, i.e., text messaging
type TextUI struct {
	in  <-chan string
	out chan<- string
	uid string
}

// Constructs a new text ui
func NewTextUI(in <-chan string, userID string) (*TextUI, <-chan string) {
	out := make(chan string)

	ui := &TextUI{
		in:  in,
		out: out,
		uid: userID,
	}

	return ui, out
}

// send is the abstraction for sending out
func (u *TextUI) send(txt string) {
	u.out <- txt
}

// Ask asks the user for input using the given query. The response is
// returned as the given string, or an error.
func (u *TextUI) Ask(s string) (string, error) {
	u.send(s)
	select {
	case msg := <-u.in:
		return msg, nil
	case <-time.After(5 * time.Minute):
		u.out <- "timeout"
		return "", fmt.Errorf("TextUI Ask, timeout")
	}
}

// AskSecret asks the user for input using the given query, but does not echo
// the keystrokes to the terminal.
func (u *TextUI) AskSecret(s string) (string, error) {
	return u.Ask(s)
}

// Output is called for normal standard output.
func (u *TextUI) Output(s string) {
	u.send(s)
}

// Info is called for information related to the previous output.
// In general this may be the exact same as Output, but this gives
// Ui implementors some flexibility with output formats.
func (u *TextUI) Info(s string) {
	u.send(s)
}

// Error is used for any error messages that might appear on standard
// error.
func (u *TextUI) Error(s string) {
	u.send(s)
}

// Warn is used for any warning message that might appear on standard
// error
func (u *TextUI) Warn(s string) {
	u.send(s)
}
