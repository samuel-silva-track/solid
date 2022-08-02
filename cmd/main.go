package main

import (
	"context"

	"github.com/samuel-silva-track/solid/internal/container"
	"github.com/samuel-silva-track/solid/pkg/domain/order"
	"github.com/samuel-silva-track/solid/pkg/domain/product"
)

func main() {

	ctx := context.Background()

	ctx, dep, err := container.New(ctx)

	//TODO: handle with panic
	if err != nil {
		panic(err)
	}

	order1 := order.Order{
		Id:         1,
		CostumerId: 1,
		OrderItems: []order.OrderItem{
			{Product: product.Product{Name: "Banana", Price: 1.3}, Quantity: 1},
			{Product: product.Product{Name: "Alface", Price: 1.3}, Quantity: 1},
		}}

	dep.Services.Order.Pay(ctx, order1)
}
