package inventory

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	productstockpb "github.com/InakiGT/processor/api-gateway/src/api/pb/product_stock/v1"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type InventoryHandler struct {
	client productstockpb.ProductStockServiceClient
}

func NewInventoryHandler(client productstockpb.ProductStockServiceClient) *InventoryHandler {
	return &InventoryHandler{client}
}

func (h *InventoryHandler) ListProductsStocks(ctx *gin.Context) {
	grpcRequest := &productstockpb.ListProductStocksRequest{}

	timeOutCtx, cancel := context.WithTimeout(ctx.Request.Context(), 10*time.Second)
	defer cancel()

	res, err := h.client.ListProductStocks(timeOutCtx, grpcRequest)
	if err != nil {
		st, _ := status.FromError(err)

		switch st.Code() {
		case codes.DeadlineExceeded:
			ctx.JSON(http.StatusGatewayTimeout, gin.H{
				"error": "product stock service timed out",
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Unknown error fetching products stocks",
			})
		}

		log.Printf("An error ocurred while trying to get products stocks: %v\n", err)

		return
	}

	ctx.JSON(http.StatusOK, toProductsStocksResponse(res.Products))
}

func (h *InventoryHandler) GetProductStock(ctx *gin.Context) {
	reqId := ctx.Param("id")

	id, err := strconv.Atoi(reqId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	grpcReq := &productstockpb.GetProductStockRequest{
		Id: uint32(id),
	}

	timeOutCtx, cancel := context.WithTimeout(ctx.Request.Context(), 10*time.Second)
	defer cancel()

	res, err := h.client.GetProductStock(timeOutCtx, grpcReq)
	if err != nil {
		st, _ := status.FromError(err)

		switch st.Code() {
		case codes.InvalidArgument:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		case codes.DeadlineExceeded:
			ctx.JSON(http.StatusGatewayTimeout, gin.H{
				"error": "product stock service timed out",
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "An unknown error ocurred while fetching product stock",
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, toProductStockResponse(res.Product))
}
