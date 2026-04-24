package price

import (
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"net/http"
	"payment-service/internal/domain"
	"payment-service/internal/domain/price"
	dtoprice "payment-service/internal/http/dto/price"
	"payment-service/internal/http/respond"
)

func (h *Handler) CreateCoursePrice(c *gin.Context) {
	var request dtoprice.CoursePriceRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		respond.JSON(c, http.StatusBadRequest, "invalid body")
		return
	}
	result, err := h.client.CreateCoursePrice(c.Request.Context(), convertCoursePriceRequest(request))
	if err != nil {
		writePriceError(c, err)
		return
	}
	respond.JSON(c, http.StatusOK, convertCoursePrice(result))
}

func (h *Handler) UpdateCoursePrice(c *gin.Context) {
	id, err := uuid.Parse(c.Param("id"))
	if err != nil {
		respond.JSON(c, http.StatusBadRequest, "invalid course id")
		return
	}
	var request dtoprice.UpdateCoursePriceRequest
	if err = c.ShouldBindJSON(&request); err != nil {
		respond.JSON(c, http.StatusBadRequest, "invalid body")
		return
	}
	result, err := h.client.UpdateCoursePrice(c.Request.Context(), id, convertUpdatePriceRequest(request))
	if err != nil {
		writePriceError(c, err)
		return
	}
	respond.JSON(c, http.StatusOK, convertCoursePrice(result))
}

func convertCoursePrice(resp *price.CoursePrice) dtoprice.CoursePriceResponse {
	return dtoprice.CoursePriceResponse{
		ID:       resp.ID,
		CourseID: resp.CourseID,
		Amount:   resp.Amount,
		Currency: resp.Currency,
	}
}

func convertCoursePriceRequest(resp dtoprice.CoursePriceRequest) *price.CoursePrice {
	return &price.CoursePrice{
		CourseID: resp.CourseID,
		Amount:   resp.Amount,
		Currency: resp.Currency,
	}
}
func convertUpdatePriceRequest(resp dtoprice.UpdateCoursePriceRequest) *price.CoursePrice {
	return &price.CoursePrice{
		Amount:   resp.Amount,
		Currency: resp.Currency,
	}
}
func writePriceError(c *gin.Context, err error) {
	switch {
	case errors.Is(err, domain.ErrValidation):
		respond.Error(c, http.StatusBadRequest, "validation", "invalid request query")
	default:
		c.Error(err)
		respond.Error(c, http.StatusInternalServerError, "internal", domain.ErrInternal.Error())
	}
}
