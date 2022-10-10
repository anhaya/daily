package main

type OrderTask struct {
	Order
}

type Order struct {
	User      string
	OrderType string
	Product   string
	Price     float64
}
