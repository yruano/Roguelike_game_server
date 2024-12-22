package server

import(
  "server/game_client"
)

// 움직임 처리
func handleMove(client *game_client.Client, move game_client.MoveData) {
	mutex.Lock()
	client.X = move.X
	client.Y = move.Y
	mutex.Unlock()

	// 모든 클라이언트에게 업데이트 전송
	broadcastPlayerStates()
}
