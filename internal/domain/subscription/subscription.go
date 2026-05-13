package subscription

import "github.com/google/uuid"

type Subscription struct {
	ID       uuid.UUID `gorm:"column:id"`
	UserID   uuid.UUID `gorm:"column:user_id"`
	CourseID uuid.UUID `gorm:"column:course_id"`
}

func (Subscription) TableName() string {
	return "course_subscription"
}
