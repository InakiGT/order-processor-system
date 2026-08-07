package changeorderstatus

import "github.com/InakiGT/processor/order-service/src/internal/domain/entities"

type ChangeOrderStatusCommand struct {
	OrderID entities.OrderID
	Status  entities.OrderStatus
}
