package hub

import (
	"wyvern-server/internal/managers/channel"
	"wyvern-server/internal/models"
)

func (m *HubManager) AddHub(h *models.Hub) {
	m.mu.Lock()
	m.Hubs[h.ID] = HubToInstance(m.Pg, h)
	m.mu.Unlock()
}

func (m *HubManager) RemoveHub(h *models.Hub) {
	m.mu.Lock()
	delete(m.Hubs, h.ID)
	m.mu.Unlock()
}

func (m *HubManager) AddChannel(h *models.Channel) {
	m.mu.Lock()
	// Gets The Hub, and addes the channel to the HubInstance
	m.Hubs[h.HubID].Channels[h.ID] = channel.ChannelToInstance(h)
	m.mu.Unlock()
}

func (m *HubManager) RemoveChannel(h *models.Channel) {
	m.mu.Lock()
	delete(m.Hubs[h.HubID].Channels, h.ID)
	m.mu.Unlock()
}
