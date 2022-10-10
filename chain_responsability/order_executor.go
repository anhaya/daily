package main

type OrderExecutor interface {
	SetNext(OrderExecutor)
	Execute(orderTask *OrderTask)
}
