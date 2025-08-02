package hub

import (
	"context"
	"encoding/json"
	"fmt"

	"wyvern-server/internal/dbctx"
	"wyvern-server/internal/managers/channel"
	"wyvern-server/internal/models"
)

func (m *HubManager) AddHub(h *models.Hub) {
	fmt.Printf("[HubManager] AddHub Called for ID: %d\n", h.ID)

	m.mu.Lock()
	m.Hubs[h.ID] = HubToInstance(m.Pg, h)
	m.mu.Unlock()
}

func (m *HubManager) RemoveHub(h *models.Hub) {
	fmt.Printf("[HubManager] RemoveHub Called for ID: %d\n", h.ID)
	m.mu.Lock()
	delete(m.Hubs, h.ID)
	m.mu.Unlock()
}

func (m *HubManager) AddChannel(h *models.Channel) {
	fmt.Printf("[HubManager] AddChannel Called for ID: %d\n", h.ID)
	m.mu.Lock()
	// Gets The Hub, and addes the channel to the HubInstance
	m.Hubs[h.HubID].Channels[h.ID] = channel.ChannelToInstance(h)
	m.mu.Unlock()
}

func (m *HubManager) RemoveChannel(h *models.Channel) {
	fmt.Printf("[HubManager] RemoveChannel Called for ID: %d\n", h.ID)

	m.mu.Lock()
	delete(m.Hubs[h.HubID].Channels, h.ID)
	m.mu.Unlock()
}

func (m *HubManager) RegisterClient(h *models.Client) {
	fmt.Printf("[HubManager] RegisterClient Called for Hub ID: %d\n", h.HubID)

	m.mu.Lock()
	defer m.mu.Unlock()

	m.Hubs[h.HubID].OnlineClients[h] = true
}

func (m *HubManager) UnregisterClient(h *models.Client) {
	fmt.Printf("[HubManager] UnregisterClient Called for Hub ID: %d\n", h.HubID)

	m.mu.Lock()
	defer m.mu.Unlock()

	delete(m.Hubs[h.HubID].OnlineClients, h)
}

func (m *HubManager) HandleIncomingMsg(h *models.Message) {
	fmt.Printf("[HubManager] UnregisterClient Called for Hub ID: %d\n", h.HubID)

	coll := dbctx.GetCtx().Mongo.Database.Collection("messages")

	coll.InsertOne(context.Background(), h)

	data, _ := json.Marshal(h)

	m.mu.RLock()
	defer m.mu.RUnlock()

	for client := range m.Hubs[h.HubID].OnlineClients {
		client.Send <- data
	}
}
