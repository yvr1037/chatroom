package model

type Message struct {
	ID        uint64 `json:"message_id" gorm:"autoIncrement"`
	CreatedAt int64  `json:"created_at" gorm:"autoCreateTime:milli"`
	UpdatedAt int64  `json:"updated_at" gorm:"autoUpdateTime:milli"`
	// UserID    uint64 `json:"user_id"`
	*User `json:"user"`
}
