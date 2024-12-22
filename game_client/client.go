package game_client

import (
	"github.com/gorilla/websocket"
)

// 클라이언트 관리
type Client struct {
	Id   string
	Conn *websocket.Conn
	X, Y float32
}
