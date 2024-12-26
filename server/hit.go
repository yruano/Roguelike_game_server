package server

import (
	"server/game_client"
)

func handleHit(client *game_client.Client, hit game_client.HitData) {
	client.PlayerState.HP -= hit.Damage

	switch hit.Debuffer_type {
	case "speed":
    client.PlayerState.Speed -= hit.Debuffer_cost
	case "defense":
    client.PlayerState.Defense -= hit.Debuffer_cost
	}

	broadcastPlayerStates()
}
