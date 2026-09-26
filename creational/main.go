package main

import (
	. "creational/abstract-factory"
	. "creational/factory-method"
	"fmt"
)

func main() {
	// Factory method
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
}
