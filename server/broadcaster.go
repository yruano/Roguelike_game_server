package server

import (
	"fmt"
	"server/game_client"
)

// 플레이어 상태 전송
func broadcastPlayerStates() {
	mutex.Lock()
	defer mutex.Unlock()

	players := make(map[string]map[string]float32)
	for id, client := range clients {
		players[id] = map[string]float32{"x": client.X, "y": client.Y}
	}

	for _, client := range clients {
		err := client.Conn.WriteJSON(players)
		if err != nil {
			fmt.Println("WriteJSON error:", err)
			client.Conn.Close()
			delete(clients, client.Id)
		}
	}
}

// 공격 정보 전송
func broadcastAttack(attack game_client.AttackData) {
	for _, client := range clients {
		err := client.Conn.WriteJSON(attack)
		if err != nil {
			fmt.Println("WriteJSON error:", err)
			client.Conn.Close()
			delete(clients, client.Id)
		}
	}
}
