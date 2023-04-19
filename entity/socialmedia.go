package entity

// SocialMedia represents the model for an social media
type SocialMedia struct {
	ID             uint64 `gorm:"primaryKey" json:"id"`
	UserID         uint64 `json:"user_id"`
	User           User   `json:"user,omitempty"`
	Name           string `gorm:"not null" json:"name"`
	SocialMediaUrl string `gorm:"not null" json:"social_media_url"`
	GORMModel
}
