package adapter

import (
	"fmt"
	"math"
	"net/http"

	"github.com/google/uuid"
)

type PaymentProcessor interface {
	ProcessPayment(amount float64, currency string) (int, string)
	RefundPayment(transactionId string) int
}

type LegacyPaySystem struct {
	transactionHistory map[string]string
}

func (l *LegacyPaySystem) makeCharge(cents int, currencyCode string) (int, string) {
	transactionId := uuid.New().String()
	amount := fmt.Sprintf("%d %s", cents, currencyCode)
	l.transactionHistory[transactionId] = amount
	fmt.Printf("A payment of %s cents successful with transaction ID %s\n", amount, transactionId)
	return http.StatusOK, transactionId
}

func (l *LegacyPaySystem) issueReversal(txnRef string) int {
	amount, ok := l.transactionHistory[txnRef]
	if !ok {
		return http.StatusNotFound
	}
	fmt.Printf("A refund of %s successfully initiated\n", amount)
	return http.StatusOK
}

type LegacyPaymentProcessor struct {
	legacyPaymentSystem *LegacyPaySystem
}

func NewLegacyPaymentProcessor() PaymentProcessor {
	return &LegacyPaymentProcessor{
		legacyPaymentSystem: &LegacyPaySystem{
			transactionHistory: make(map[string]string),
		},
	}
}

func (l *LegacyPaymentProcessor) ProcessPayment(amount float64, currency string) (int, string) {
	cents := int(math.Round(amount * 100))
	return l.legacyPaymentSystem.makeCharge(cents, currency)
}

func (l *LegacyPaymentProcessor) RefundPayment(transactionId string) int {
	return l.legacyPaymentSystem.issueReversal(transactionId)
}
