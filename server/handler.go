package server

import (
	"encoding/json"
	"fmt"
	"net/http"
	"server/game_client"
	"sync"

	"github.com/gorilla/websocket"
)

var clients = make(map[string]*game_client.Client)
var mutex = &sync.Mutex{}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// handleConnections 함수 수정
func HandleConnections(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		fmt.Println("Upgrade error:", err)
		return
	}

	clientID := fmt.Sprintf("%d", len(clients)+1)
	client := &game_client.Client{Id: clientID, Conn: conn, X: 0, Y: 0}
	mutex.Lock()
	clients[clientID] = client
	mutex.Unlock()

	defer func() {
		mutex.Lock()
		delete(clients, clientID)
		mutex.Unlock()
		conn.Close()
	}()

	for {
		// 클라이언트에서 메시지 배열 수신
		var actions []game_client.ClientAction
		err := conn.ReadJSON(&actions)
		if err != nil {
			fmt.Println("ReadJSON error:", err)
			break
		}

		// 각각의 동작을 처리
		for _, msg := range actions {
			switch msg.Action {
			case "move":
				var move game_client.MoveData
				err := json.Unmarshal(msg.Data, &move)
				if err != nil {
					fmt.Println("Invalid move data:", err)
					continue
				}
				handleMove(client, move)

			case "attack":
				var attack game_client.AttackData
				err := json.Unmarshal(msg.Data, &attack)
				if err != nil {
					fmt.Println("Invalid attack data:", err)
					continue
				}
				handleAttack(client, attack)

			default:
				fmt.Println("Unknown action:", msg.Action)
			}
		}
	}
}
