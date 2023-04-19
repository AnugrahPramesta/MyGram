package entity

// Photo represents the model for an photo
type Photo struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	UserID   uint64 `json:"user_id"`
	User     User   `json:"user,omitempty"`
	Title    string `gorm:"not null" json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `gorm:"not null" json:"photo_url"`
	GORMModel
}
