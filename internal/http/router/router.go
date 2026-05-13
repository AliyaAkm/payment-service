package router

import (
	"github.com/gin-gonic/gin"
	"payment-service/internal/http/handlers/epayment"
	"payment-service/internal/http/handlers/order"
	"payment-service/internal/http/handlers/paymentmethod"
	"payment-service/internal/http/handlers/price"
)

type Handler struct {
	Order         *order.Handler
	Price         *price.Handler
	PaymentMethod *paymentmethod.Handler
	Payment       *epayment.Handler
}

func New(handler Handler, globalMiddlewares []gin.HandlerFunc) *gin.Engine {
	r := gin.New()
	r.Use(globalMiddlewares...)

	r.GET("/health", health)

	// order
	r.POST("/order", handler.Order.CreateOrder)

	// course price
	r.POST("/price", handler.Price.CreateCoursePrice)
	r.PUT("/price/:id", handler.Price.UpdateCoursePrice)

	r.GET("/payment_method", handler.PaymentMethod.GetAllPaymentMethod)

	// payment
	r.POST("/payment/token", handler.Payment.GetToken)
	r.POST("/payment", handler.Payment.CreatePayment)
	r.GET("/payment/transaction/:invoice_id", handler.Payment.GetTransaction)
	return r
}
