package model

type User struct {
	ID        uint64 `json:"user_id" gorm:"autoIncrement"`
	Username  string `json:"user_name"`
	NickName  string `json:"nick_name"`
	Password  string `json:"password"`
	CreatedAt int    `json:"created_at"`
	UpdatedAt int    `json:"updated_at"`
}
