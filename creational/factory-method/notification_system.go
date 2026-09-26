package factorymethod

import (
	"errors"
	"fmt"
)

type Platform string

const (
	ANDROID Platform = "android"
	IOS     Platform = "ios"
	WEB     Platform = "web"
)

type Notification interface {
	Send(string)
}

type smsNotification struct{}

func (s smsNotification) Send(msg string) {
	fmt.Printf("[Android SMS] %s\n", msg)
}

type pushNotification struct{}

func (s pushNotification) Send(msg string) {
	fmt.Printf("[iOS Push] %s\n", msg)
}

type emailNotification struct{}

func (s emailNotification) Send(msg string) {
	fmt.Printf("[Web Email] %s\n", msg)
}

func GetNotification(plat Platform) (Notification, error) {
	switch plat {
	case ANDROID:
		return smsNotification{}, nil
	case IOS:
		return pushNotification{}, nil
	case WEB:
		return emailNotification{}, nil
	}
	return nil, errors.New("invalid platform type")
}
