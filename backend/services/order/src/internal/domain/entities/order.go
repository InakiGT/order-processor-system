package entities

import (
	"time"

	"github.com/InakiGT/processor/order-service/src/internal/domain/errors"
)

type OrderID int
type OrderStatus string

const (
	StatusPending   OrderStatus = "PENDING"
	StatusPaid      OrderStatus = "PAID"
	StatusCancelled OrderStatus = "CANCELLED"
	StatusFailed    OrderStatus = "FAILED"
)

type Order struct {
	OrderId     OrderID
	Status      OrderStatus
	TotalAmount int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Items       []*OrderItem
}

func NewOrder(items []*OrderItem) (*Order, error) {
	if len(items) == 0 {
		return nil, errors.ErrEmptyInventory
	}

	order := &Order{Status: StatusPending, Items: items}
	order.CalcTotal()

	return order, nil
}

func (p *Order) CalcTotal() int32 {
	var total int32

	for _, item := range p.Items {
		total += item.PriceAtPurchase
	}

	return total
}

func (o *Order) ChangeStatus(status OrderStatus) {
	o.Status = status
	o.UpdatedAt = time.Now()
}
