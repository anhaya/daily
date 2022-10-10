package main

import "fmt"

type OrderUser struct {
	next OrderExecutor
}

func (o *OrderUser) Execute(orderTask *OrderTask) {
	if orderTask.User == "Carlos Anhaya" {
		o.next.Execute(orderTask)
		return
	}

	fmt.Println("Invalid Name")
}

func (o *OrderUser) SetNext(orderExecutor OrderExecutor) {
	o.next = orderExecutor
}
