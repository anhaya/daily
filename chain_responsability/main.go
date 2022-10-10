package main

import "fmt"

func main() {
	fmt.Println("this func")
	orderDomain := &OrderTask{
		Order: Order{
			User:      "Carlos Anhaya",
			OrderType: "SELL",
			Product:   "BTC",
			Price:     23.000,
		},
	}

	orderPrice := &OrderPrice{}

	orderProduct := &OrderProduct{}
	orderProduct.SetNext(orderPrice)

	orderType := &OrderType{}
	orderType.SetNext(orderPrice)

	orderUser := &OrderUser{}
	orderUser.SetNext(orderType)

	orderUser.Execute(orderDomain)
}
