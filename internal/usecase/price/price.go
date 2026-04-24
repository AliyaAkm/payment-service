package price

import (
	"context"
	"github.com/google/uuid"
	"payment-service/internal/domain/price"
)

func (u *UseCase) CreateCoursePrice(ctx context.Context, value *price.CoursePrice) (*price.CoursePrice, error) {
	value.ID = uuid.New()

	err := u.repo.CreateCoursePrice(ctx, value)
	if err != nil {
		return nil, err
	}
	return u.repo.GetCoursePriceByID(ctx, value.ID)
}

func (u *UseCase) UpdateCoursePrice(ctx context.Context, id uuid.UUID, newValue *price.CoursePrice) (*price.CoursePrice, error) {
	oldValue, err := u.repo.GetCoursePriceByID(ctx, id)
	if err != nil {
		return nil, err
	}

	if newValue.Amount != 0 {
		oldValue.Amount = newValue.Amount
	}
	if newValue.Currency != "" {
		oldValue.Currency = newValue.Currency
	}

	if err = u.repo.UpdateCoursePrice(ctx, id, oldValue); err != nil {
		return nil, err
	}

	return u.repo.GetCoursePriceByID(ctx, id)
}
