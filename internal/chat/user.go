package chat

import (
	"chatroom/internal/model"
	"errors"
	"io"
	"time"

	"golang.org/x/net/context"
	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

type User struct {
	UserID   uint64 `json:"user_id"`
	NickName string `json:"nick_name"`
	EnterAt  int64  `json:"-"`

	messageChannel chan *Message
	conn           *websocket.Conn
}

var System = &User{UserID: 0, NickName: "System"}

func NewUser(u *model.User, conn *websocket.Conn) *User {
	return &User{
		UserID:         u.ID,
		NickName:       u.NickName,
		EnterAt:        time.Now().Unix(),
		messageChannel: make(chan *Message, 32),
		conn:           conn,
	}
}

func (u *User) Send(msg *Message) {
	u.messageChannel <- msg
}

func (u *User) SendMessage(ctx context.Context) {
	for msg := range u.messageChannel {
		wsjson.Write(ctx,u.conn,msg)
	}
}

func (u *User) CloseMessageChannel() {
	close(u.messageChannel)
}

func (u *User) ReceiveMessage(ctx context.Context) error {
	var (
		// receiveMessage string
		receiveMessage map[string]string
		err            error
	)
	for {
		err = wsjson.Read(ctx, u.conn, &receiveMessage)
		if err != nil {
			var closeErr websocket.CloseError
			if errors.As(err, &closeErr) {
				return nil
			} else if errors.Is(err, io.EOF) {
				return nil
			}

			return err
		}

		sendMessage := NewMessage(u,receiveMessage["message_content"])
		Broadcaster.Broadcast(sendMessage)
	}
}
