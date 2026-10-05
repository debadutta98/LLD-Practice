package bridge

import (
	"fmt"
	"strings"
)

type MessageChannel interface {
	SendRaw(header string, body string)
}

type EmailChannel struct{}

func (e *EmailChannel) SendRaw(header string, body string) {
	fmt.Printf("[EMAIL] Header: %s | Body: %s\n", header, body)
}

type SMSChannel struct{}

func (s *SMSChannel) SendRaw(header string, body string) {
	fmt.Printf("[SMS] %s - %s\n", header, body)
}

type PushChannel struct{}

func (p *PushChannel) SendRaw(header string, body string) {
	fmt.Printf("[PUSH NOTIFICATION] %s: %s\n", header, body)
}

type Priority string

const (
	HIGH   Priority = "HIGH"
	MEDIUM Priority = "MEDIUM"
	LOW    Priority = "LOW"
)

type Notification struct {
	channel MessageChannel
}

type AlertNotification struct {
	Notification
	Priority Priority
	Content  string
}

func NewAlertNotification(channel MessageChannel, priority Priority, content string) *AlertNotification {
	return &AlertNotification{
		Notification: Notification{channel: channel},
		Priority:     priority,
		Content:      content,
	}
}

func (a *AlertNotification) Notify() {
	header := fmt.Sprintf("ALERT [%s]", a.Priority)
	a.channel.SendRaw(header, a.Content)
}

type DigestNotification struct {
	Notification
	Messages []string
}

func NewDigestNotification(channel MessageChannel, messages []string) *DigestNotification {
	return &DigestNotification{
		Notification: Notification{channel: channel},
		Messages:     messages,
	}
}

func (d *DigestNotification) Notify() {
	header := "Daily Digest"
	body := strings.Join(d.Messages, " | ")
	d.channel.SendRaw(header, body)
}
