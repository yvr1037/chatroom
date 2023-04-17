package chat

import (
	"time"
)

type Message struct {
	User     *User  `json:"user"`
	Content  string `json:"message_content`
	SendTime int64  `json:"send_time"`
}

func NewMessage(user *User, content string) *Message {
	message := &Message{
		User:     user,
		Content:  content,
		SendTime: time.Now().Unix(),
	}
	return message
}

func NewUserEnterMessage(user *User) *Message {
	return &Message{
		User:     System,
		Content:  "welcome" + user.NickName + "enter chatroom",
		SendTime: time.Now().Unix(),
	}
}

func NewUserLeaveMessage(user *User) *Message {
	return &Message{
		User:System,
		Content: user.NickName + "leave chatroom",
		SendTime: time.Now().Unix(),
	}
}

func NewErrorMessage(content string) *Message {
	return &Message{
		User:System,
		Content: content,
		SendTime: time.Now().Unix(),
	}
}
