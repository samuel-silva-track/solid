package order

import (
	"context"
	"fmt"
)

type ServiceI interface {
	Pay(ctx context.Context, order Order)
}

type Service struct {
	Repository RepositoryI
}

func NewService(repository RepositoryI) (*Service, error) {
	if repository == nil {
		return nil, ErrEmptyRepository
	}

	return &Service{
		Repository: repository,
	}, nil
}

func (o *Service) Pay(ctx context.Context, order Order) {
	var total float64
	for _, item := range order.OrderItems {
		total += item.Product.Price * float64(item.Quantity)

	}

	fmt.Println(total)
	fmt.Println(order.PaymentMethod)
}
