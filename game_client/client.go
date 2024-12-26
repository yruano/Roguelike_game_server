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
	PlayerState PlayerState // 플레이어의 상태를 관리하는 구조체
}

type PlayerState struct {
	MaxHp   float32
	MaxMp   float32
	HP      float32
	MP      float32
	Speed   float32
	Defense float32
	Level   int
	XP      float32
}
