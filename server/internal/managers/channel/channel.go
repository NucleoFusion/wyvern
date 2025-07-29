package channel

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	"wyvern-server/internal/models"
)

func NewChannelManager(pg *sql.DB) *ChannelManager {
	fmt.Println("[Managers] Creating Channel Manager....")
	res, err := pg.Query("SELECT id, hub_id, name, private FROM channel")
	if err != nil {
		log.Fatal(err.Error())
	}

	items := make(map[int]*ChannelInstance, 0)
	for res.Next() {
		var ch models.Channel
		err := res.Scan(&ch.ID, &ch.HubID, &ch.Name, &ch.Private)
		if err != nil {
			log.Fatal(err.Error())
		}

		items[ch.ID] = ChannelToInstance(&ch)
	}

	fmt.Printf("[Managers] Successfully Created Hub Manager w/ Entries: %d\n", len(items))

	return &ChannelManager{
		Channels: items,
		mu:       sync.RWMutex{},
	}
}

func ChannelToInstance(ch *models.Channel) *ChannelInstance {
	return &ChannelInstance{
		ID:      ch.ID,
		Clients: make(map[*models.Client]bool),
		mu:      sync.Mutex{},
	}
}
