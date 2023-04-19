package dto

import (
	"project-mygram/entity"
	"time"
)

type SocialMediaCreateReq struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
}

type SocialMediaListRes struct {
	ResponseList
	Data []entity.SocialMedia `json:"data"`
}

type SocialMediaUpdateReq struct {
	Name           string `json:"name" validate:"required"`
	SocialMediaUrl string `json:"social_media_url" validate:"required"`
}

type SocialMediaCreateUpdateResponse struct {
	ID             uint64    `json:"id"`
	Name           string    `json:"name" `
	SocialMediaUrl string    `json:"social_media_url" `
	UserID         uint64    `json:"user_id"`
	CreatedAt      time.Time `json:"created_at,omitempty"`
	UpdatedAt      time.Time `json:"updated_at,omitempty"`
}
