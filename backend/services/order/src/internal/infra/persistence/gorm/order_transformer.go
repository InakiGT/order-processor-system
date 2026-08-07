package persistence

import "github.com/InakiGT/processor/order-service/src/internal/domain/entities"

func toOrderEntities(orders []Order) []*entities.Order {
	newOrders := make([]*entities.Order, 0, len(orders))

	for _, order := range orders {
		newOrders = append(newOrders, toOrderEntity(order))
	}

	return newOrders
}

func toOrderEntity(order Order) *entities.Order {
	return &entities.Order{
		OrderId:     entities.OrderID(order.ID),
		Status:      order.Status,
		TotalAmount: order.TotalAmount,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
		Items:       toItemEntities(order.Items),
	}
}

func toOrderModel(order *entities.Order) *Order {
	return &Order{
		Status:      order.Status,
		TotalAmount: order.TotalAmount,
		CreatedAt:   order.CreatedAt,
		UpdatedAt:   order.UpdatedAt,
		Items:       toItemModels(order.Items),
	}
}
