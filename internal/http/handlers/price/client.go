package price

import (
	"context"
	"github.com/google/uuid"
	"payment-service/internal/domain/price"
)

type client interface {
	UpdateCoursePrice(ctx context.Context, id uuid.UUID, newValue *price.CoursePrice) (*price.CoursePrice, error)
	CreateCoursePrice(ctx context.Context, value *price.CoursePrice) (*price.CoursePrice, error)
}
