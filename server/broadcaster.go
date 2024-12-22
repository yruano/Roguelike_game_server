package server

import (
	"fmt"

	"server/game_client"
)

// 모든 클라이언트에게 플레이어 상태를 브로드캐스트하는 함수
func broadcastPlayerStates() {
	mutex.Lock()
	defer mutex.Unlock()

	players := make(map[string]map[string]interface{})
	for id, client := range clients {
		players[id] = map[string]interface{}{
			"x":     client.X,
			"y":     client.Y,
			"hp":    client.PlayerStats.HP,
			"mp":    client.PlayerStats.MP,
			"level": client.PlayerStats.Level,
			"xp":    client.PlayerStats.XP,
		}
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
