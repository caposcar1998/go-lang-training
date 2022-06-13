package main

import (
	"errors"
	"fmt"
	"time"
)

type Order struct {
	Id          string
	orderLines  []OrderLine
	total       float32
	dateCreated time.Time
	user        string
}

type OrderLine struct {
	Id          string
	item        string
	dateCreated time.Time
	quantity    int16
	unitPrice   float32
}

func CreateOrder(id string, orderLine []OrderLine, dateCreated time.Time, user string) (*Order, error) {

	var total float32 = 0

	if id == "" {
		return nil, errors.New("Some values are missing")
	} else if orderLine == nil {
		return nil, errors.New("Some values are missing")
	} else if dateCreated.IsZero() {
		return nil, errors.New("Some values are missing")
	} else if user == "" {
		return nil, errors.New("Some values are missing")
	} else {
		total = calculateCost(orderLine)
	}

	return &Order{
		Id:          id,
		orderLines:  orderLine,
		total:       total,
		dateCreated: dateCreated,
		user:        user,
	}, nil

}

func AppendOrderLine(orderLine OrderLine, order *Order) {
	order.orderLines = append(order.orderLines, orderLine)
	order.total = calculateCost(order.orderLines)
}

func UpdateOrderLineItem(order *Order, item string, idOrderLineEdit string) {
	for index, element := range order.orderLines {
		if element.Id == idOrderLineEdit {
			order.orderLines[index].item = item
		}
	}
	order.total = calculateCost(order.orderLines)

}

func UpdateOrderLineQuantity(order *Order, quantity int16, idOrderLineEdit string) {
	for index, element := range order.orderLines {
		if element.Id == idOrderLineEdit {
			order.orderLines[index].quantity = quantity

		}
	}
	order.total = calculateCost(order.orderLines)

}

func UpdateOrderLinePrice(order *Order, unitPrice float32, idOrderLineEdit string) {
	for index, element := range order.orderLines {
		if element.Id == idOrderLineEdit {
			order.orderLines[index].unitPrice = unitPrice
		}
	}
	order.total = calculateCost(order.orderLines)

}

func RemoveOrderLine(order *Order, idRemove string) {
	for index, element := range order.orderLines {
		if element.Id == idRemove {
			order.orderLines = append(order.orderLines[:index], order.orderLines[index+1:]...)
		}
	}
	order.total = calculateCost(order.orderLines)
}

func PrintPrice(order *Order) float32 {
	return order.total
}

func calculateCost(orderLines []OrderLine) float32 {
	var finalTotal float32 = 0

	for _, element := range orderLines {
		finalTotal += (float32(element.quantity) * float32(element.unitPrice))
	}

	return finalTotal
}

func main() {

	currentTime := time.Now()
	var orders = []OrderLine{{"0", "pizza", currentTime, 10, 40.0}}

	// Create new order
	fmt.Println("CREATE NEW ORDER")
	order, error := CreateOrder("1", orders, currentTime, "oscar")

	if error == nil {
		fmt.Println(order)
		fmt.Println()
		//Add order item
		fmt.Println("ADD ORDER ITEM")
		var orderItem OrderLine = OrderLine{"1", "tacos", currentTime, 5, 10.0}
		AppendOrderLine(orderItem, order)
		fmt.Println(order)
		fmt.Println()

		//Remove order item
		fmt.Println("REMOVE ORDER ITEM")
		RemoveOrderLine(order, "1")
		fmt.Println(order)
		fmt.Println()

		//Update order item
		fmt.Println("UPDATE ORDER ITEM NAME")
		UpdateOrderLineItem(order, "chilaquiles", "0")
		fmt.Println(order)
		fmt.Println()

		//Update order item price
		fmt.Println("UPDATE ORDER ITEM PRICE")
		UpdateOrderLinePrice(order, 50.0, "0")
		fmt.Println(order)
		fmt.Println()

		//Update order item quantity
		fmt.Println("UPDATE ORDER ITEM QUANTITY")
		UpdateOrderLineQuantity(order, 20, "0")
		fmt.Println()

		//print price
		fmt.Println("PRINT FINAL PRICE")
		PrintPrice(order)
		fmt.Println("The final price is: ", order.total)
		fmt.Println()

	} else {
		fmt.Println(error)
	}

}
