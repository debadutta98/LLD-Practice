package main

import (
	. "structural/adapter"
	. "structural/bridge"
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
}
