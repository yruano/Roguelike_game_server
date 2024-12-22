package server

import (
  "server/game_client"
  "fmt"
)

// 공격 처리
func handleAttack(client *game_client.Client, attack game_client.AttackData) {
	fmt.Printf("Player %s attacked at (%f, %f) with angle %.2f\n", attack.PlayerID, attack.X, attack.Y, attack.Angle)

	// 공격 정보를 모든 클라이언트에게 전송
	broadcastAttack(attack)
}

