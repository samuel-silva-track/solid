package container

import (
	"context"

	"github.com/samuel-silva-track/solid/pkg/domain/order"
)

//log, queues, cache, etc.
type components struct {
}

// serviço com regras de negocio
type services struct {
	Order order.ServiceI
}

type Dependecy struct {
	Components components
	Services   services
}

func New(ctx context.Context) (context.Context, *Dependecy, error) {
	orderRepository := order.NewRepository()

	orderService, err := order.NewService(orderRepository)
	if err != nil {
		return nil, nil, err
	}

	svr := services{
		Order: orderService,
	}

	dep := Dependecy{
		Components: components{},
		Services:   svr,
	}

	return ctx, &dep, nil
}
