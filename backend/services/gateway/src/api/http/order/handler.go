package order

import (
	"net/http"

	pb "github.com/InakiGT/processor/api-gateway/src/api/pb/order/v1"
	"github.com/gin-gonic/gin"
)

type OrderHandler struct {
	orderClient pb.OrderServiceClient
}

func NewOrderHandler(client pb.OrderServiceClient) *OrderHandler {
	return &OrderHandler{client}
}

func (h *OrderHandler) CreateOrder(ctx *gin.Context) {
	var req CreateOrderRequest

	if err := ctx.ShouldBindBodyWithJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid request",
		})
	}

	grpcReq := &pb.CreateOrderRequest{
		Items: toCreateOrderGrpcItems(req.items),
	}

	res, err := h.orderClient.CreateOrder(ctx.Request.Context(), grpcReq)
	if err != nil {
		return
	}

	ctx.JSON(http.StatusCreated, toCreateOrderJSONResponse(res))
}
