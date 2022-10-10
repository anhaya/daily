package main

import "fmt"

type OrderProduct struct {
	next OrderExecutor
}

func (o *OrderProduct) Execute(orderTask *OrderTask) {
	if orderTask.Order.Product == "BTC" {
		o.next.Execute(orderTask)
		return
	}

	fmt.Println("Invalid Order Product")
}

func (o *OrderProduct) SetNext(orderExecutor OrderExecutor) {
	o.next = orderExecutor
}
