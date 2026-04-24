package order

import (
	"context"
	"github.com/google/uuid"
	"payment-service/internal/domain/order"
)

const pendingStatus = "pending"

func (u *UseCase) CreateOrder(ctx context.Context, value *order.Order) (*order.Order, error) {
	value.ID = uuid.New()

	status, err := u.statusRepo.GetOrderStatusByCode(ctx, pendingStatus)
	if err != nil {
		return nil, err
	}
	value.StatusID = status.ID

	price, err := u.priceRepo.GetCoursePriceByCourseID(ctx, value.CourseID)
	if err != nil {
		return nil, err
	}
	value.Amount = price.Amount
	value.Currency = price.Currency

	err = u.repo.CreateOrder(ctx, value)
	if err != nil {
		return nil, err
	}
	return u.repo.GetOrderByID(ctx, value.ID)
}
