package usecase

import (
	"github.com/danielzinhors/clean-go-expert/internal/entity"
)

type OrdersUseCase struct {
	OrderRepository entity.OrderRepositoryInterface
}

type OrdersOutputDTO struct {
	ID         string  `json:"id"`
	Price      float64 `json:"price"`
	Tax        float64 `json:"tax"`
	FinalPrice float64 `json:"final_price"`
}

func NewListOrdersUseCase(
	orderRepository entity.OrderRepositoryInterface,
) *OrdersUseCase {
	return &OrdersUseCase{
		OrderRepository: orderRepository,
	}
}

func (c *OrdersUseCase) Execute() ([]OrdersOutputDTO, error) {
	orders, err := c.OrderRepository.FindAll()
	if err != nil {
		return []OrdersOutputDTO{}, err
	}

	var ordersOutput []OrdersOutputDTO
	for _, order := range orders {
		ordersOutput = append(ordersOutput, OrdersOutputDTO{
			ID:         order.ID,
			Price:      order.Price,
			Tax:        order.Tax,
			FinalPrice: order.FinalPrice,
		})
	}

	return ordersOutput, nil
}
