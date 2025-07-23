package channel

import (
	"sync"

	"wyvern-server/internal/models"
)

func ChannelToInstance(ch *models.Channel) *ChannelInstance {
	return &ChannelInstance{
		ID:      ch.ID,
		Clients: make(map[*models.Client]bool),
		mu:      sync.Mutex{},
	}
}
