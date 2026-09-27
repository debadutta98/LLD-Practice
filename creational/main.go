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

	orderA := NewMealBuilder(MainItem{Name: "Burger", Protein: "Chicken", Price: 8.99}).
		WithSide(SideItem{Name: "Fries", Size: Large, Price: 3.49}).
		WithDrink(Drink{Name: "Soda", Size: Medium, HasIce: true, Price: 2.10}).
		Build()

	orderB := NewMealBuilder(MainItem{Name: "Wrap", Protein: "Veggie", Price: 7.50}).
		WithDessert(Dessert{Name: "Cookie", Price: 1.50}).
		Build()

	orderA.DisplayMealSummary()
	orderB.DisplayMealSummary()
}
