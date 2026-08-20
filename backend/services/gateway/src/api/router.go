package api

import (
	"github.com/InakiGT/processor/api-gateway/src/api/http/inventory"
	"github.com/InakiGT/processor/api-gateway/src/api/http/order"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	orderHandler *order.OrderHandler,
	inventoryHandler *inventory.InventoryHandler,
) *gin.Engine {
	router := gin.Default()

	orders := router.Group("/orders")
	{
		orders.GET("", orderHandler.ListOrders)
		orders.GET("/:id", orderHandler.GetOrder)
		orders.POST("", orderHandler.CreateOrder)
		orders.DELETE("/:id", orderHandler.Delete)
		orders.PATCH("/", orderHandler.ChangeOrderStatus)
	}

	inventory := router.Group("/inventory/items")
	{
		inventory.PATCH("/:id", inventoryHandler.Restock)
		inventory.GET("", inventoryHandler.ListProductsStocks)
		inventory.GET("/:id", inventoryHandler.GetProductStock)
		inventory.POST("", inventoryHandler.CreateProductStock)
		inventory.DELETE("", inventoryHandler.DeleteProductStock)
	}

	return router
}
