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
			"x":       client.X,
			"y":       client.Y,
			"maxhp":   client.PlayerState.MaxHp,
			"maxmp":   client.PlayerState.MaxMp,
			"hp":      client.PlayerState.HP,
			"mp":      client.PlayerState.MP,
			"speed":   client.PlayerState.Speed,
			"defense": client.PlayerState.Defense,
			"level":   client.PlayerState.Level,
			"xp":      client.PlayerState.XP,
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
