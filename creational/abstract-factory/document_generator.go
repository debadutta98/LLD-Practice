package abstractfactory

import "fmt"

type Document interface {
	Print()
}

type Invoice interface {
	Document
}

type TaxSummary interface {
	Document
}

type Market interface {
	GetInvoice() Invoice
	GetTaxSummary() TaxSummary
}

type usInvoice struct {
	price int
}

func (u *usInvoice) Print() {
	fmt.Printf("US Invoice: Billed in USD ($%d). Remit payment to US Domestic Treasury Account.\n", u.price)
}

type usTaxSummary struct{}

func (u *usTaxSummary) Print() {
	fmt.Println("US Tax Summary: Breakdown of State Sales Tax (e.g., CA: 7.25%, NY: 4.00%).")
}

type euInvoice struct {
	price int
}

func (e *euInvoice) Print() {
	fmt.Printf("EU Invoice: Billed in EUR (€%d). Includes mandatory cross-border VAT note (Reverse Charge applicable).\n", e.price)
}

type euTaxSummary struct{}

func (e *euTaxSummary) Print() {
	fmt.Println("EU Tax Summary: Breakdown of Country VAT rates (e.g., DE: 19%, FR: 20%). Compliant with GDPR data handling.")
}

type USMarket struct {
	DefaultPrice int
}

func NewUSMarket(price int) Market {
	return &USMarket{DefaultPrice: price}
}

func (u *USMarket) GetInvoice() Invoice {
	return &usInvoice{price: u.DefaultPrice}
}

func (u *USMarket) GetTaxSummary() TaxSummary {
	return &usTaxSummary{}
}

type EUMarket struct {
	DefaultPrice int
}

func NewEUMarket(price int) Market {
	return &EUMarket{DefaultPrice: price}
}

func (e *EUMarket) GetInvoice() Invoice {
	return &euInvoice{price: e.DefaultPrice}
}

func (e *EUMarket) GetTaxSummary() TaxSummary {
	return &euTaxSummary{}
}
