package decorator

import (
	"encoding/base64"
	"fmt"
	"strings"
	"time"
)

type Notifer interface {
	Send(message string)
}

type EmailNotifier struct {
	recepients []string
}

func NewEmailNotifier(recepients ...string) *EmailNotifier {
	return &EmailNotifier{recepients}
}

func (e *EmailNotifier) Send(msg string) {
	fmt.Println("[EMAIL] To:", strings.Join(e.recepients, ", "), msg)
}

type SlackDecorator struct {
	notifier Notifer
	channels []string
}

func NewSlackDecorator(channels ...string) *SlackDecorator {
	return &SlackDecorator{
		channels: channels,
	}
}

func (s *SlackDecorator) SetNotifer(notifer Notifer) {
	s.notifier = notifer
}

func (s *SlackDecorator) Send(msg string) {
	if s.notifier != nil {
		s.notifier.Send(msg)
	}
	fmt.Println("[Slack] Channels:", strings.Join(s.channels, ","), msg)
}

type SMSDecorator struct {
	contact_numbers []string
	notifer         Notifer
}

func NewSMSDecorator(contact_numbers []string) *SMSDecorator {
	return &SMSDecorator{
		contact_numbers: contact_numbers,
	}
}

func (s *SMSDecorator) SetNotifer(notifier Notifer) {
	s.notifer = notifier
}

func (s *SMSDecorator) Send(msg string) {
	if s.notifer != nil {
		s.notifer.Send(msg)
	}
	fmt.Println("[SMS] TO:", strings.Join(s.contact_numbers, ", "), msg)
}

type TimestampDecorator struct {
	notifier Notifer
}

func NewTimestampDecorator(notifer Notifer) Notifer {
	return &TimestampDecorator{notifer}
}

func (t *TimestampDecorator) Send(msg string) {
	t.notifier.Send(fmt.Sprintf("[%s] %s", time.Now().Format(time.RFC3339), msg))
}

type EncryptionDecorator struct {
	notifer Notifer
}

func NewEncryptionDecorator(notifer Notifer) Notifer {
	return &EncryptionDecorator{notifer}
}

func (t *EncryptionDecorator) Send(msg string) {
	t.notifer.Send(base64.RawStdEncoding.EncodeToString([]byte(msg)))
}
