package hub

import (
	"database/sql"
	"log"
	"sync"

	"wyvern-server/internal/managers/channel"
	"wyvern-server/internal/models"
)

// Finds all Hub Entries in DB and adds them to manager
func NewHubManager(pg *sql.DB) *HubManager {
	res, err := pg.Query("SELECT id, repo, owner_id FROM hub")
	if err != nil {
		log.Fatal(err.Error())
	}

	items := make(map[int]*HubInstance, 0)
	for res.Next() {
		var hub models.Hub
		err := res.Scan(&hub.ID, &hub.Repo, &hub.OwnerID)
		if err != nil {
			log.Fatal(err.Error())
		}

		items[hub.ID] = HubToInstance(pg, &hub)
	}

	return &HubManager{
		Hubs: items,
		mu:   sync.RWMutex{},
	}
}

// Converting Hub to HubInstance (Will also query for all channels in hub)
func HubToInstance(pg *sql.DB, hub *models.Hub) *HubInstance {
	res, err := pg.Query("SELECT id, name, private FROM channel WHERE hub_id = $1", hub.ID)
	if err != nil {
		log.Fatal(err.Error())
	}

	items := make(map[int]*channel.ChannelInstance, 0)
	for res.Next() {
		var ch models.Channel
		err := res.Scan(&ch.ID, &ch.Name, &ch.Private)
		if err != nil {
			log.Fatal(err.Error())
		}
		ch.HubID = hub.ID

		items[ch.ID] = channel.ChannelToInstance(&ch)
	}

	return &HubInstance{
		ID:            hub.ID,
		OwnerID:       hub.OwnerID,
		Repo:          hub.Repo,
		OnlineClients: make(map[*models.Client]bool),
		Channels:      items,
		mu:            sync.Mutex{},
	}
}
