package price

import "github.com/google/uuid"

type CoursePriceRequest struct {
	CourseID uuid.UUID `json:"course_id"`
	Amount   int       `json:"amount"`
	Currency string    `json:"currency"`
}

type UpdateCoursePriceRequest struct {
	Amount   int    `json:"amount"`
	Currency string `json:"currency"`
}

type CoursePriceResponse struct {
	ID       uuid.UUID `json:"id"`
	CourseID uuid.UUID `json:"course_id"`
	Amount   int       `json:"amount"`
	Currency string    `json:"currency"`
}
