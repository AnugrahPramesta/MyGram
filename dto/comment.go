package dto

import (
	"project-mygram/entity"
	"time"
)

type CommentCreateUpdateReq struct {
	PhotoID uint64 `json:"photo_id"`
	Message string `json:"message" validate:"required"`
}

type CommentListRes struct {
	ResponseList
	Data []entity.Comment `json:"data"`
}

type CommentCreateUpdateResponse struct {
	ID        uint64    `json:"id"`
	Message   string    `json:"message"`
	PhotoID   uint64    `json:"photo_id"`
	UserID    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
