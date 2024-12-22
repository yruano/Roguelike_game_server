package game_client

import (
	"github.com/gorilla/websocket"
)

// Client 구조체 정의
type Client struct {
	Id          string
	Conn        *websocket.Conn
	X           float32
	Y           float32
	PlayerStats PlayerStats // 플레이어의 상태를 관리하는 구조체
}

type PlayerStats struct {
	HP    int
	MP    int
	Level int
	XP    int
}
