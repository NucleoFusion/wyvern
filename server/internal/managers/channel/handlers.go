package channel

import (
	"fmt"

	"wyvern-server/internal/models"
)

func (m *ChannelManager) AddChannel(h *models.Channel) {
	fmt.Printf("[ChannelManager] AddChannel Called for ID: %d\n", h.ID)

	m.mu.Lock()
	// Gets The Hub, and addes the channel to the HubInstance
	m.Channels[h.ID] = ChannelToInstance(h)
	m.mu.Unlock()
}

func (m *ChannelManager) RemoveChannel(h *models.Channel) {
	fmt.Printf("[ChannelManager] RemoveChannel Called for ID: %d\n", h.ID)

	m.mu.Lock()
	delete(m.Channels, h.ID)
	m.mu.Unlock()
}
