package order

import "github.com/samuel-silva-track/solid/pkg/domain/product"

type Order struct {
	Id            int
	CostumerId    int
	OrderItems    []OrderItem
	PaymentMethod string
}

type OrderItem struct {
	Product  product.Product
	Quantity int
}
