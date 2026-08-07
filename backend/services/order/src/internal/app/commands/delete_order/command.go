package deleteorder

import "github.com/InakiGT/processor/order-service/src/internal/domain/entities"

type DeleteOrderCommand struct {
	Id entities.OrderID
}
