package persistence

import "github.com/InakiGT/processor/order-service/src/internal/domain/entities"

func toItemEntities(orderItems []OrderItem) []*entities.OrderItem {
	items := make([]*entities.OrderItem, 0, len(orderItems))

	for _, item := range orderItems {
		items = append(items, toItemEntity(item))
	}

	return items
}

func toItemEntity(item OrderItem) *entities.OrderItem {
	return &entities.OrderItem{
		OrderItemID:     entities.OrderItemID(item.ID),
		OrderID:         entities.OrderID(item.OrderID),
		ProductID:       item.ProductID,
		Quantity:        item.Quantity,
		PriceAtPurchase: item.PriceAtPurchase,
	}
}

func toItemModel(item *entities.OrderItem) OrderItem {
	return OrderItem{
		OrderID:         uint(item.OrderID),
		ProductID:       item.ProductID,
		Quantity:        item.Quantity,
		PriceAtPurchase: item.PriceAtPurchase,
	}
}

func toItemModels(items []*entities.OrderItem) []OrderItem {
	newItems := make([]OrderItem, 0, len(items))

	for _, item := range items {
		newItems = append(newItems, toItemModel(item))
	}

	return newItems
}
