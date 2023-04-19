package entity

// Comment represents the model for an comment
type Comment struct {
	ID      uint64        `gorm:"primaryKey" json:"id"`
	UserID  uint64        `json:"user_id"`
	User    UserComment   `json:"user,omitempty"`
	PhotoID uint64        `json:"photo_id"`
	Photo   *PhotoComment `json:"photo,omitempty"`
	Message string        `gorm:"not null" json:"message"`
	GORMModel
}

type UserComment struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	Username string `gorm:"not null;uniqueIndex" json:"username" `
	Email    string `gorm:"not null;uniqueIndex" json:"email" `
	Age      uint64 `gorm:"not null" json:"age"`
	GORMModel
}
type PhotoComment struct {
	ID       uint64 `gorm:"primaryKey" json:"id"`
	UserID   uint64 `json:"user_id"`
	Title    string `gorm:"not null" json:"title"`
	Caption  string `json:"caption"`
	PhotoUrl string `gorm:"not null" json:"photo_url"`
	GORMModel
}

func (PhotoComment) TableName() string {
	return "photos"
}

func (UserComment) TableName() string {
	return "users"
}
