package epayment

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"payment-service/internal/domain/order"
	"payment-service/internal/domain/orderstatus"
	dtoepayment "payment-service/internal/http/dto/epayment"
	dtoorder "payment-service/internal/http/dto/order"
	"payment-service/internal/http/respond"
	"payment-service/service/epayment"
)

func (h *Handler) GetToken(c *gin.Context) {
	request := dtoepayment.GetTokenRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		respond.JSON(c, http.StatusBadRequest, "invalid body")
		return
	}
	result, err := h.client.GetToken(c.Request.Context(), convertTokenRequest(request))
	if err != nil {
		respond.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}
	respond.JSON(c, http.StatusOK, convertToken(result))
}

func (h *Handler) CreatePayment(c *gin.Context) {
	request := dtoepayment.PaymentRequest{}
	if err := c.ShouldBindJSON(&request); err != nil {
		respond.JSON(c, http.StatusBadRequest, "invalid body")
		return
	}
	result, err := h.paymentUseCase.CreatePayment(c.Request.Context(), convertPaymentRequest(request))
	if err != nil {
		respond.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}
	respond.JSON(c, http.StatusOK, convertPayment(result))
}

func (h *Handler) GetTransaction(c *gin.Context) {
	invoiceID := c.Param("invoice_id")
	result, err := h.client.GetStatusTransaction(c.Request.Context(), invoiceID)
	if err != nil {
		respond.JSON(c, http.StatusInternalServerError, err.Error())
		return
	}
	respond.JSON(c, http.StatusOK, convertTransaction(result))
}

func convertTokenRequest(req dtoepayment.GetTokenRequest) epayment.GetTokenRequest {
	return epayment.GetTokenRequest{
		Amount:   req.Amount,
		Currency: req.Currency,
	}
}

func convertPaymentRequest(req dtoepayment.PaymentRequest) epayment.PaymentRequest {
	return epayment.PaymentRequest{
		OrderID:    req.OrderID,
		Name:       req.Name,
		Cryptogram: req.Cryptogram,
		CardSave:   req.CardSave,
	}
}

func convertToken(req epayment.GetTokenResponse) dtoepayment.GetTokenResponse {
	return dtoepayment.GetTokenResponse{
		AccessToken: req.AccessToken,
		ExpiresIn:   req.ExpiresIn,
		Scope:       req.Scope,
		TokenType:   req.TokenType,
	}
}
func convertPayment(resp *order.Order) dtoorder.OrderResponse {
	return dtoorder.OrderResponse{
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

func convertTransaction(req epayment.GetTransactionResponse) dtoepayment.GetTransactionResponse {
	resp := dtoepayment.GetTransactionResponse{
		ResultCode:    req.ResultCode,
		ResultMessage: req.ResultMessage,
	}

	if req.Transaction == nil {
		resp.Transaction = nil
		return resp
	}

	resp.Transaction = &dtoepayment.Transaction{
		ID:          req.Transaction.ID,
		CreatedDate: req.Transaction.CreatedDate,
		InvoiceID:   req.Transaction.InvoiceID,

		Amount:       req.Transaction.Amount,
		AmountBonus:  req.Transaction.AmountBonus,
		PayoutAmount: req.Transaction.PayoutAmount,
		OrgAmount:    req.Transaction.OrgAmount,

		ApprovalCode: req.Transaction.ApprovalCode,
		Data:         req.Transaction.Data,

		Currency:    req.Transaction.Currency,
		Terminal:    req.Transaction.Terminal,
		TerminalID:  req.Transaction.TerminalID,
		AccountID:   req.Transaction.AccountID,
		Description: req.Transaction.Description,
		Language:    req.Transaction.Language,

		CardMask: req.Transaction.CardMask,
		CardType: req.Transaction.CardType,
		Issuer:   req.Transaction.Issuer,

		Reference:    req.Transaction.Reference,
		Reason:       req.Transaction.Reason,
		ReasonCode:   req.Transaction.ReasonCode,
		IntReference: req.Transaction.IntReference,

		Secure:        req.Transaction.Secure,
		SecureDetails: req.Transaction.SecureDetails,

		StatusID:   req.Transaction.StatusID,
		StatusName: req.Transaction.StatusName,

		Name:  req.Transaction.Name,
		Email: req.Transaction.Email,
		Phone: req.Transaction.Phone,

		CardID: req.Transaction.CardID,
		XLSRRN: req.Transaction.XLSRRN,

		IP:          req.Transaction.IP,
		IPCountry:   req.Transaction.IPCountry,
		IPCity:      req.Transaction.IPCity,
		IPRegion:    req.Transaction.IPRegion,
		IPDistrict:  req.Transaction.IPDistrict,
		IPLatitude:  req.Transaction.IPLatitude,
		IPLongitude: req.Transaction.IPLongitude,
	}

	return resp
}
