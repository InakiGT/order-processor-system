package api

import (
	"github.com/InakiGT/processor/api-gateway/src/api/http/inventory"
	"github.com/InakiGT/processor/api-gateway/src/api/http/order"
	"github.com/InakiGT/processor/api-gateway/src/internal/auth"
	"github.com/gin-gonic/gin"
)

func NewRouter(
	orderHandler *order.OrderHandler,
	inventoryHandler *inventory.InventoryHandler,
	authMiddleware gin.HandlerFunc,
) *gin.Engine {
	router := gin.Default()

	protected := router.Group("/")
	protected.Use(authMiddleware)

	orders := protected.Group("/orders")
	{
		orders.GET("", orderHandler.ListOrders)
		orders.GET("/:id", orderHandler.GetOrder)
		orders.POST("", orderHandler.CreateOrder)
		orders.DELETE("/:id", orderHandler.Delete)
		orders.PATCH("/", orderHandler.ChangeOrderStatus)
	}

	inventory := protected.Group("/inventory/items")
	{
		inventory.PATCH(
			"/:id",
			auth.RequireScope("write:inventory"),
			inventoryHandler.Restock,
		)
		inventory.GET(
			"",
			auth.RequireScope("read:inventory"),
			inventoryHandler.ListProductsStocks,
		)
		inventory.GET(
			"/:id",
			auth.RequireScope("read:inventory"),
			inventoryHandler.GetProductStock,
		)
		inventory.POST(
			"",
			auth.RequireScope("write:inventory"),
			inventoryHandler.CreateProductStock,
		)
		inventory.DELETE(
			"",
			auth.RequireScope("write:inventory"),
			inventoryHandler.DeleteProductStock,
		)
	}

	return router
}
