package model

import (
	"github.com/google/uuid"
	"time"
)

type Filter struct {
	UserID      *uuid.UUID `json:"user_id,omitempty"`
	ServiceName *string    `json:"service_name,omitempty"`
	Price       *int       `json:"price,omitempty"`
	StartDate   *time.Time `json:"start_date,omitempty"`
	EndDate     *time.Time `json:"end_date,omitempty"`
	Limit       *int       `json:"limit,omitempty"`
	Offset      *int       `json:"offset,omitempty"`
}

func (f *Filter) Normalize() {

	const (
		defaultLimit  = 10
		defaultOffset = 0
		maxLimit      = 50
	)

	if f.Limit == nil || *f.Limit <= 0 || *f.Limit > maxLimit {
		l := defaultLimit
		f.Limit = &l
	}

	if f.Offset == nil || *f.Offset < 0 {
		o := defaultOffset
		f.Offset = &o
	}

}
