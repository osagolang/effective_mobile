package model

import (
	"github.com/google/uuid"
)

// Потенциально можно вынести в отдельный пакет реквестов (когда их будет много)
type CreateSubscriptionRequest struct {
	ServiceName string    `json:"service_name" binding:"required"`
	Price       int       `json:"price" binding:"required"`
	UserID      uuid.UUID `json:"user_id" binding:"required"`
	StartDate   string    `json:"start_date" binding:"required"`
}

type UpdateSubscriptionRequest struct {
	ServiceName *string `json:"service_name,omitempty"`
	Price       *int    `json:"price,omitempty"`
	EndDate     *string `json:"end_date,omitempty"`
}
