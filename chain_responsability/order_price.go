package main

import "fmt"

type OrderPrice struct {
	next OrderExecutor
}

func (o *OrderPrice) Execute(orderTask *OrderTask) {
	if orderTask.Price > 1 {
		return
	}

	fmt.Println("Invalid Order Price")
}

func (o *OrderPrice) SetNext(orderExecutor OrderExecutor) {
	o.next = orderExecutor
}
