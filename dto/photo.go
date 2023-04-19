package dto

import (
	"project-mygram/entity"
	"time"
)

type PhotoCreateUpdateReq struct {
	Title    string `json:"title" validate:"required"`
	PhotoUrl string `json:"photo_url" validate:"required"`
	Caption  string `json:"caption"`
}

type PhotoListRes struct {
	ResponseList
	Data []entity.Photo `json:"data"`
}

type PhotoCreateUpdateResponse struct {
	ID        uint64    `json:"id"`
	Title     string    `json:"title"`
	PhotoUrl  string    `json:"photo_url"`
	Caption   string    `json:"caption"`
	UserID    uint64    `json:"user_id"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}
