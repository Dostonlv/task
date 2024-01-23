package models

import (
	"time"

	"github.com/google/uuid"
)

// Blog model
type Blog struct {
	ID        uuid.UUID `json:"id" db:"id" validate:"omitempty,uuid"`
	Title     string    `json:"title" db:"title" validate:"required,gte=3"`
	Content   string    `json:"content" db:"content" validate:"required,gte=15"`
	CreatedAt time.Time `json:"created_at" db:"created_at"`
}

// BlogList model
type BlogList struct {
	TotalCount int     `json:"total_count"`
	TotalPages int     `json:"total_pages"`
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	HasMore    bool    `json:"has_more"`
	Blogs      []*Blog `json:"blogs"`
}

type BlogSwagger struct {
	Title   string `json:"title" validate:"required,gte=3"`
	Content string `json:"content" validate:"required,gte=15"`
}
