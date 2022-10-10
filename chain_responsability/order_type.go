package main

import "fmt"

type OrderType struct {
	next OrderExecutor
}

func (o *OrderType) Execute(orderTask *OrderTask) {
	if orderTask.OrderType == "SELL" {
		o.next.Execute(orderTask)
		return
	}

	fmt.Println("Invalid Order Type")
}

func (o *OrderType) SetNext(orderExecutor OrderExecutor) {
	o.next = orderExecutor
}
