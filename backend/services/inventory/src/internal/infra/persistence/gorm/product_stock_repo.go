package gorm

import (
	"context"
	"time"

	"github.com/InakiGT/processor/inventory-service/src/internal/domain/entities"
	"gorm.io/gorm"
)

type ProductStockGormRepo struct {
	db *gorm.DB
}

type ProductStock struct {
	gorm.Model
	SKU               entities.SKU
	Brand             string
	ModelName         string
	AvailableQuantity int
	ReservedQuantity  int
	UpdatedAt         time.Time
	CreatedAt         time.Time
}

func NewProductStockRepository(db *gorm.DB) *ProductStockGormRepo {
	return &ProductStockGormRepo{db}
}

func (r *ProductStockGormRepo) FindAll(ctx context.Context) ([]*entities.ProductStock, error) {
	var productStocks []*ProductStock
	if err := r.db.WithContext(ctx).Find(&productStocks).Error; err != nil {
		return nil, err
	}

	return toProductStockEntities(productStocks), nil
}

func (r *ProductStockGormRepo) FindOneByID(ctx context.Context, id entities.ProductID) (*entities.ProductStock, error) {
	var productStock *ProductStock
	if err := r.db.WithContext(ctx).Where("id = ?", id).First(&productStock).Error; err != nil {
		return nil, err
	}

	return toProductStockEntity(productStock), nil
}

func (r *ProductStockGormRepo) Save(ctx context.Context, productStock *entities.ProductStock) (*entities.ProductStock, error) {
	gormProduct := toProductStockModel(productStock)
	if err := r.db.WithContext(ctx).Save(gormProduct).Error; err != nil {
		return nil, err
	}

	return toProductStockEntity(gormProduct), nil
}

func (r *ProductStockGormRepo) Update(ctx context.Context, productStock *entities.ProductStock) error {
	var productToUpdate ProductStock

	if err := r.db.WithContext(ctx).Where("id = ?", productStock.ProductID).First(&productToUpdate).Error; err != nil {
		return err
	}

	productToUpdate = *toProductStockModel(productStock)
	return r.db.WithContext(ctx).Save(&productToUpdate).Error
}

func (r *ProductStockGormRepo) Delete(ctx context.Context, id entities.ProductID) error {
	return r.db.WithContext(ctx).Where("id = ?", id).Delete(&ProductStock{}).Error
}
