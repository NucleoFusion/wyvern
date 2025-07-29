package hub

import (
	"fmt"

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
