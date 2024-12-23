package server

import (
	"server/game_client"
)

func handleHit(client *game_client.Client, hit game_client.HitData) {
	client.PlayerStats.HP -= hit.Damage

	switch hit.Debuffer_type {
	case "speed":
    client.PlayerStats.Speed -= hit.Debuffer_cost
	case "defense":
    client.PlayerStats.Defense -= hit.Debuffer_cost
	}

	broadcastPlayerStates()
}
