package game_client

import (
	"encoding/json"
)

// MoveData 구조체
type MoveData struct {
	X float32 `json:"x"`
	Y float32 `json:"y"`
}

// HitData 구조체
type HitData struct {
  Damage float32 `json:"damage"`
  Debuffer_type string `json:"debuffer_type"`
  Debuffer_cost float32 `json:"debuffer_cost"`
}

// AttackData 구조체
type AttackData struct {
	PlayerID string  `json:"player_id"`
	Type     string  `json:"type"`
	X        float32 `json:"x"`
	Y        float32 `json:"y"`
	Angle    float64 `json:"angle"`
  Cost     float32 `json:"cost"`
}

// 클라이언트가 보내는 메시지 타입
type ClientAction struct {
	Action string          `json:"action"` // "move" 또는 "attack"
	Data   json.RawMessage `json:"data"`   // 행동 데이터 (JSON으로 처리)
}
