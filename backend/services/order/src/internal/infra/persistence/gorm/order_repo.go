package persistence

import (
	"context"
	"time"

	"github.com/InakiGT/processor/order-service/src/internal/domain/entities"
	"gorm.io/gorm"
)

type OrderRepository struct {
	db *gorm.DB
}

type Order struct {
	gorm.Model
	TotalAmount int32
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Items       []OrderItem
	Status      entities.OrderStatus
}

type OrderItem struct {
	gorm.Model
	OrderID         uint
	ProductID       int
	Quantity        int
	PriceAtPurchase int32
}

func NewOrderRepository(db *gorm.DB) *OrderRepository {
	return &OrderRepository{db}
}

func (r *OrderRepository) FindAll(ctx context.Context) ([]*entities.Order, error) {
	var orders []Order

	if err := r.db.WithContext(ctx).Preload("Items").Find(&orders).Error; err != nil {
		return nil, err
	}

	return toOrderEntities(orders), nil
}

func (r *OrderRepository) FindOneById(ctx context.Context, id entities.OrderID) (*entities.Order, error) {
	var order Order

	if err := r.db.WithContext(ctx).Preload("Items").Where("id = ?").First(&order).Error; err != nil {
		return nil, err
	}

	return toOrderEntity(order), nil
}

func (r *OrderRepository) Save(ctx context.Context, order *entities.Order) (*entities.Order, error) {
	var sOrder *Order

	if order.OrderId > 0 {
		if err := r.db.WithContext(ctx).Where("id = ?", order.OrderId).First(&sOrder).Error; err != nil {
			return nil, err
		}
	} else {
		sOrder = toOrderModel(order)
	}

	if err := r.db.WithContext(ctx).Save(&sOrder).Error; err != nil {
		return nil, err
	}

	return order, nil
}

func (r *OrderRepository) Delete(ctx context.Context, id entities.OrderID) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&Order{}).Error
}
