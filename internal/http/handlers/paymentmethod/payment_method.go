package paymentmethod

import (
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
	"payment-service/internal/domain"
	domainpaymentmethod "payment-service/internal/domain/paymentmethod"
	"payment-service/internal/http/dto/paymentmethod"
	"payment-service/internal/http/respond"
)

func (h *Handler) GetAllPaymentMethod(c *gin.Context) {
	result, err := h.client.GetAllPaymentMethod(c.Request.Context())
	if err != nil {
		writeCatalogError(c, err)
		return
	}
	respond.JSON(c, http.StatusOK, convertProviders(result))
}

func convertProviders(resp []domainpaymentmethod.PaymentMethod) []paymentmethod.PaymentMethod {
	providers := make([]paymentmethod.PaymentMethod, len(resp))

	for i := range resp {
		providers[i] = paymentmethod.PaymentMethod{
			ID:   resp[i].ID,
			Name: resp[i].Name,
			Code: resp[i].Code,
		}
	}
	return providers
}

func writeCatalogError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrValidation):
		respond.Error(c, http.StatusBadRequest, "validation", "invalid request query")
	default:
		c.Error(err)
		respond.Error(c, http.StatusInternalServerError, "internal", domain.ErrInternal.Error())
	}
}
