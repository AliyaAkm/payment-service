package router

import (
	"github.com/gin-gonic/gin"
	"payment-service/internal/http/handlers/order"
	"payment-service/internal/http/handlers/price"
)

type Handler struct {
	Order *order.Handler
	Price *price.Handler
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
	return r
}
