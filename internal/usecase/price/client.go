package price

import (
	"context"
	"github.com/google/uuid"
	"payment-service/internal/domain/price"
)

type Repository interface {
	CreateCoursePrice(ctx context.Context, value *price.CoursePrice) error
	GetCoursePriceByCourseID(ctx context.Context, courseID uuid.UUID) (*price.CoursePrice, error)
	UpdateCoursePrice(ctx context.Context, id uuid.UUID, value *price.CoursePrice) error
	GetCoursePriceByID(ctx context.Context, id uuid.UUID) (*price.CoursePrice, error)
}
