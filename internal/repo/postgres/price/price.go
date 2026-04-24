package price

import (
	"context"
	"github.com/google/uuid"
	"payment-service/internal/domain/price"
)

func (r *Repo) CreateCoursePrice(ctx context.Context, value *price.CoursePrice) error {
	err := r.db.WithContext(ctx).Create(value).Error
	if err != nil {
		return err
	}
	return nil
}
func (r *Repo) GetCoursePriceByCourseID(ctx context.Context, courseID uuid.UUID) (*price.CoursePrice, error) {
	var entity *price.CoursePrice
	err := r.db.WithContext(ctx).Where("course_id = ?", courseID).First(&entity).Error
	if err != nil {
		return nil, err
	}
	return entity, nil
}

func (r *Repo) UpdateCoursePrice(ctx context.Context, id uuid.UUID, value *price.CoursePrice) error {
	err := r.db.WithContext(ctx).Where("id = ?", id).Updates(value).Error
	if err != nil {
		return err
	}
	return nil
}

func (r *Repo) GetCoursePriceByID(ctx context.Context, id uuid.UUID) (*price.CoursePrice, error) {
	var entity price.CoursePrice

	err := r.db.WithContext(ctx).
		Where("id = ?", id).
		First(&entity).Error
	if err != nil {
		return nil, err
	}

	return &entity, nil
}
