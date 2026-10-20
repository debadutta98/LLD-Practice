package main

import (
	"fmt"
	. "structural/adapter"
	. "structural/bridge"
	. "structural/decorator"
)

func main() {
	legacyPaymentProcessor := NewLegacyPaymentProcessor()
	_, txnId := legacyPaymentProcessor.ProcessPayment(100, "$")
	legacyPaymentProcessor.RefundPayment(txnId)

	email := &EmailChannel{}
	sms := &SMSChannel{}

	urgentAlert := NewAlertNotification(sms, HIGH, "Database connection failed!")
	urgentAlert.Notify()

	dailySummary := NewDigestNotification(email, []string{
		"User A registered",
		"Server rebooted successfully",
		"Backup completed",
	})

	dailySummary.Notify()

	// Decorator

	expresso := NewEspresso()
	milkExpresso := NewMilkDecorator(expresso)
	mochaMilkExpresso := NewMochaDecorator(milkExpresso)
	whipMochaMilkExpresso := NewWhipDecorator(mochaMilkExpresso)

	fmt.Println(whipMochaMilkExpresso.GetDescription(), whipMochaMilkExpresso.GetCost())

	emailNotifer := NewEmailNotifier("abc@email.com")
	slackNotifer := NewSlackDecorator("#ask-me-anything")
	slackNotifer.SetNotifer(emailNotifer)
	//encryptionDecorator := NewEncryptionDecorator(slackNotifer)
	timestampDecorator := NewTimestampDecorator(slackNotifer)
	timestampDecorator.Send("Hello! How are you ? :>")
}
