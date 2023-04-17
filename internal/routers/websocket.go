package routers

import (
	"chatroom/internal/chat"
	"chatroom/internal/service"
	"log"

	"github.com/gin-gonic/gin"
	"nhooyr.io/websocket"
)

func WebsocketHandler(c *gin.Context) {
	conn, err := websocket.Accept(c.Writer, c.Request, &websocket.AcceptOptions{
		InsecureSkipVerify: true,
	})
	if err != nil {
		log.Println("websocket accept error:", err)
		return
	}
	// wsjson.Write(c, conn, "test")

	ID, exists := c.Get("UserID")
	if !exists {
		log.Println("jwt middleware error")
	}
	svc := service.New(c.Request.Context())
	userGot, err := svc.UserGet(ID.(uint64))
	user := chat.NewUser(userGot, conn)

	go user.SendMessage(c.Request.Context())

	chat.Broadcaster.UserEntering(user)
	msg := chat.NewUserEnterMessage(user)
	chat.Broadcaster.Broadcast(msg)
	log.Println("user:", user.NickName, "joins chat")

	err = user.ReceiveMessage(c.Request.Context())

	chat.Broadcaster.UserLeaving(user)
	msg = chat.NewUserLeaveMessage(user)
	chat.Broadcaster.Broadcast(msg)
	log.Println("user:", user.NickName, "leaves chat")

	if err == nil {
		conn.Close(websocket.StatusNormalClosure, "")
	} else {
		log.Println("read from client error:", err)
		conn.Close(websocket.StatusInternalError, "Read from client error")
	}
}
