package channel

import (
	"wyvern-server/internal/models"
)

func (m *ChannelManager) AddChannel(h *models.Channel) {
	m.mu.Lock()
	// Gets The Hub, and addes the channel to the HubInstance
	m.Channels[h.ID] = ChannelToInstance(h)
	m.mu.Unlock()
}

func (m *ChannelManager) RemoveChannel(h *models.Channel) {
	m.mu.Lock()
	delete(m.Channels, h.ID)
	m.mu.Unlock()
}
