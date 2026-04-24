package order

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"payment-service/internal/domain"
	domainorder "payment-service/internal/domain/order"
	"payment-service/internal/domain/orderstatus"
	"payment-service/internal/http/dto/order"
	"payment-service/internal/http/respond"
)

func (h *Handler) CreateOrder(c *gin.Context) {
	request := order.OrderRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		respond.JSON(c, http.StatusBadRequest, "invalid body")
		return
	}
	result, err := h.client.CreateOrder(c.Request.Context(), convertOrderRequest(request))
	if err != nil {
		writeOrderError(c, err)
		return
	}
	respond.JSON(c, http.StatusOK, convertOrder(result))
}

func convertOrder(resp *domainorder.Order) order.OrderResponse {

	return order.OrderResponse{
		ID:       resp.ID,
		UserID:   resp.UserID,
		CourseID: resp.CourseID,
		Amount:   resp.Amount,
		Currency: resp.Currency,
		Status: orderstatus.OrderStatus{
			ID:   resp.Status.ID,
			Name: resp.Status.Name,
			Code: resp.Status.Code,
		},
	}
}

func convertOrderRequest(resp order.OrderRequest) *domainorder.Order {
	return &domainorder.Order{
		UserID:   resp.UserID,
		CourseID: resp.CourseID,
	}
}

func writeOrderError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrValidation):
		respond.Error(c, http.StatusBadRequest, "validation", "invalid request query")
	default:
		c.Error(err)
		respond.Error(c, http.StatusInternalServerError, "internal", domain.ErrInternal.Error())
	}
}
