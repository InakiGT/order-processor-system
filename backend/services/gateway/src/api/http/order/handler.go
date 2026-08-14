package order

import (
	"context"
	"log"
	"net/http"
	"strconv"
	"time"

	pb "github.com/InakiGT/processor/api-gateway/src/api/pb/order/v1"
	"github.com/gin-gonic/gin"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
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
			"error": err.Error(),
		})

		return
	}

	grpcReq := &pb.CreateOrderRequest{
		Items: toCreateOrderGrpcItems(req.Items),
	}

	timeOutCtx, cancel := context.WithTimeout(ctx.Request.Context(), 10*time.Second)
	defer cancel()

	res, err := h.orderClient.CreateOrder(timeOutCtx, grpcReq)
	if err != nil {
		st, _ := status.FromError(err)

		switch st.Code() {
		case codes.InvalidArgument:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": st.Message(),
			})
		case codes.DeadlineExceeded:
			ctx.JSON(http.StatusGatewayTimeout, gin.H{
				"error": "Order service timed out",
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Unknown error creating the order",
			})
		}

		log.Printf("Error while trying to create the order: %v \n", err)
		return
	}

	ctx.JSON(http.StatusCreated, toCreateOrderJSONResponse(res))
}

func (h *OrderHandler) ListOrders(ctx *gin.Context) {
	grpcReq := &pb.ListOrdersRequest{}

	timeOutCtx, cancel := context.WithTimeout(ctx.Request.Context(), 10*time.Second)
	defer cancel()

	res, err := h.orderClient.ListOrders(timeOutCtx, grpcReq)
	if err != nil {
		st, _ := status.FromError(err)

		switch st.Code() {
		case codes.InvalidArgument:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		case codes.DeadlineExceeded:
			ctx.JSON(http.StatusGatewayTimeout, gin.H{
				"error": "Order service timed out",
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Unknown error fetching the orders",
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, toListOrdersResponse(res))
}

func (h *OrderHandler) GetOrder(ctx *gin.Context) {
	reqId := ctx.Param("id")

	id, err := strconv.Atoi(reqId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request: Id must be a number",
		})

		return
	}

	grpcReq := &pb.GetOrderRequest{
		Id: int32(id),
	}

	timeOutCtx, cancel := context.WithTimeout(ctx.Request.Context(), 10*time.Second)
	defer cancel()

	res, err := h.orderClient.GetOrder(timeOutCtx, grpcReq)
	if err != nil {
		st, _ := status.FromError(err)

		switch st.Code() {
		case codes.InvalidArgument:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		case codes.DeadlineExceeded:
			ctx.JSON(http.StatusGatewayTimeout, gin.H{
				"error": "Order service timed out",
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Error while trying to get the order",
			})
		}

		return
	}

	ctx.JSON(http.StatusOK, toGetOrderJSONResponse(res.Order))
}

func (h *OrderHandler) Delete(ctx *gin.Context) {
	reqId := ctx.Param("id")

	id, err := strconv.Atoi(reqId)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": "Id must be a number",
		})
	}

	grpcReq := &pb.DeleteOrderRequest{
		Id: int32(id),
	}

	timeOutCtx, cancel := context.WithTimeout(ctx.Request.Context(), 10*time.Second)
	defer cancel()

	if _, err := h.orderClient.DeleteOrder(timeOutCtx, grpcReq); err != nil {
		st, _ := status.FromError(err)

		switch st.Code() {
		case codes.InvalidArgument:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		case codes.DeadlineExceeded:
			ctx.JSON(http.StatusGatewayTimeout, gin.H{
				"error": "Order service timed out",
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "error trying to delete the order",
			})
		}

		log.Printf("Error while trying to delete an order: %v", err)
		return
	}

	ctx.Status(http.StatusNoContent)
}

func (h *OrderHandler) ChangeOrderStatus(ctx *gin.Context) {
	var req ChangeOrderStatusRequest

	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})

		return
	}

	grpcReq := &pb.ChangeOrderStatusRequest{
		Id:     req.OrderId,
		Status: req.Status,
	}

	timeOutCtx, cancel := context.WithTimeout(ctx.Request.Context(), 10*time.Second)
	defer cancel()

	if _, err := h.orderClient.ChangeOrderStatus(timeOutCtx, grpcReq); err != nil {
		st, _ := status.FromError(err)

		switch st.Code() {
		case codes.InvalidArgument:
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
		case codes.DeadlineExceeded:
			ctx.JSON(http.StatusGatewayTimeout, gin.H{
				"error": "Order service timed out",
			})
		default:
			ctx.JSON(http.StatusInternalServerError, gin.H{
				"error": "Unknown error ocurred while changing order status",
			})
		}

		log.Printf("Error while trying to change order status: %v\n", err)
		return
	}

	ctx.Status(http.StatusNoContent)
}
