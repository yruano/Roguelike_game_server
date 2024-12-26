package server

import (
	"math/rand"
	"time"

	"server/game_client"
)

// client에서 플레이 캐릭터를 선택하면 스탯을 세팅함
func player_init(client *game_client.Client, state game_client.PlayerState) {
	client.PlayerState = state
	broadcastPlayerStates()
}

func player_spawn() {
	mutex.Lock()
	defer mutex.Unlock()

	// 로컬 랜덤 생성기 초기화
	rnd := rand.New(rand.NewSource(time.Now().UnixNano()))

	for _, client := range clients {
		// 랜덤 x, y 좌표 생성 (예: 0~100의 범위)
		client.X = float32(rnd.Intn(101))
		client.Y = float32(rnd.Intn(101))
	}

	broadcastPlayerStates()
}
