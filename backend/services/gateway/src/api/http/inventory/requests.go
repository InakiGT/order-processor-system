package inventory

type CreateProductStockRequest struct {
	Brand          string `json:"brand" binding:"required"`
	Model          string `json:"model" binding:"required"`
	AvailableStock int32  `json:"available_stock" binding:"required"`
}

type RestockRequest struct {
	Id       uint32 `json:"product_stock_id" binding:"required"`
	Quantity uint32 `json:"increase" binding:"required"`
}
