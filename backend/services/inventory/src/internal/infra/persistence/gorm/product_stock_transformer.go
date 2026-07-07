package gorm

import "github.com/InakiGT/processor/inventory-service/src/internal/domain/entities"

func toProductStockEntity(model *ProductStock) *entities.ProductStock {
	return &entities.ProductStock{
		ProductID:         entities.ProductID(model.ID),
		SKU:               model.SKU,
		Model:             model.ModelName,
		Brand:             model.Brand,
		AvailableQuantity: model.AvailableQuantity,
		ReservedQuantity:  model.ReservedQuantity,
		UpdatedAt:         model.UpdatedAt,
		CreatedAt:         model.CreatedAt,
	}
}

func toProductStockEntities(models []*ProductStock) []*entities.ProductStock {
	entities := make([]*entities.ProductStock, 0)

	for _, model := range models {
		entities = append(entities, toProductStockEntity(model))
	}

	return entities
}

func toProductStockModel(entity *entities.ProductStock) *ProductStock {
	return &ProductStock{
		SKU:               entity.SKU,
		ModelName:         entity.Model,
		Brand:             entity.Brand,
		ReservedQuantity:  entity.ReservedQuantity,
		AvailableQuantity: entity.AvailableQuantity,
	}
}
