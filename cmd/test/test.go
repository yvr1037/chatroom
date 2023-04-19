package main

import (
	"context"
	"fmt"
	"time"

	"chatroom/global"
	"chatroom/internal/setting"

	"nhooyr.io/websocket"
	"nhooyr.io/websocket/wsjson"
)

func main() {
	setupSettings()
	/* 	token, err := auth.GenerateToken(4)
	   	if err != nil {
	   		panic(err)
	   	}
	   	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjozLCJpc3MiOiJIdWFuZyIsImV4cCI6MTY0ODAyODI2NCwibmJmIjoxNjQ4MDI4MjY0LCJpYXQiOjE2NDgwMjgyNjR9.Hn7VBIEOi02Hhe7uCxD4AoMzeCckizQ9SjgIL7EiUD4"
	   	claims, err := auth.ParseToken(token)
	   	if err != nil {
	   		panic(err)
	   	}
	   	fmt.Println("UserID", claims.UserID) */
	// ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	ctx, cancel := context.WithTimeout(context.Background(), time.Hour)
	defer cancel()

	// c, _, err := websocket.Dial(ctx, "ws://localhost:8080/ws", nil)
	c, _, err := websocket.Dial(ctx,
		"ws://175.178.43.145:4001/ws?token=eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VyX2lkIjoyLCJpc3MiOiJIdWFuZyIsImV4cCI6MTY0ODEwNDEzMSwibmJmIjoxNjQ4MDk2OTMxLCJpYXQiOjE2NDgwOTY5MzF9.pvLgLeNv_jaP9nu7HV03CFcm-tGrsPwi5q9d1HyR0U4",
		nil)
	if err != nil {
		panic(err)
	}
	defer c.Close(websocket.StatusInternalError, "internal error")

	err = wsjson.Write(ctx, c, "Hello WebSocker Server")
	if err != nil {
		panic(err)
	}

	var v interface{}
	err = wsjson.Read(ctx, c, &v)
	if err != nil {
		panic(err)
	}
	fmt.Printf("receive server respond: %v\n", v)

}
func setupSettings() error {
	settings, err := setting.NewSetting()
	if err != nil {
		return err
	}
	err = settings.ReadSection("Server", &global.ServerSettings)
	if err != nil {
		return err
	}
	err = settings.ReadSection("Database", &global.DatabaseSettings)
	if err != nil {
		return err
	}
	err = settings.ReadSection("JWT", &global.JWTSettings)
	if err != nil {
		return err
	}

	global.ServerSettings.ReadTimeout *= time.Second
	global.ServerSettings.WriteTimeout *= time.Second
	global.JWTSettings.Expire *= time.Second
	return nil
}
