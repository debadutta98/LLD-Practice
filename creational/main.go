package main

import (
	. "creational/abstract-factory"
	. "creational/builder"
	. "creational/factory-method"
	"fmt"
)

func main() {
	// Factory method
	fmt.Println("-----Factory method----")
	notification, err := GetNotification(ANDROID)
	if err != nil {
		panic(err)
	}
	notification.Send("Your verification code is 1234")

	converter, err := GetConverter(HTML, PDF)
	if err != nil {
		panic(err)
	}
	fmt.Println(converter.Convert("Hello World!"))
	// Abstract factory
	fmt.Println("-----Abstract factory----")
	market := NewEUMarket(10)
	market.GetInvoice().Print()
	market.GetTaxSummary().Print()

	zwave, err := GetHardWareFamily("zwave")
	doorlock := zwave.GetSmartDoorLock()
	doorlock.Lock()
	if doorlock.GetLockState() != true {
		panic("door must be locked")
	} else {
		fmt.Println("Door is locked")
	}

	// Builder

	fmt.Println("-----Builder----")

	builder := NewReportBuilder()

	reporter := NewReporter(builder)

	if report, err := reporter.BuildExecutiveSummaryReport(
		"A Memorable Press Visit to Times of India",
		"Introduction:",
		"Thanks",
		[]string{
			"The press visit commenced with a warm welcome from the Times of India's editorial team",
			"who graciously guided us through their state-of-the-art newsroom",
		},
	); err != nil {
		panic(err)
	} else {
		fmt.Print(report.Title())
	}
}
