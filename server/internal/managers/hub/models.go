package hub

import (
	"sync"

	"wyvern-server/internal/managers/channel"
	"wyvern-server/internal/models"
)

type HubMsgType int

const (
	AddClient HubMsgType = iota
	RemoveClient
	AddHub
	RemoveHub
	AddChannel
	RemoveChannel
)

type HubInstance struct {
	ID            int
	Repo          string
	OwnerID       int
	Channels      map[int]*channel.ChannelInstance // ChannelID -> ChannelInstance
	OnlineClients map[*models.Client]bool          // Just use map for 0(1) access
	mu            sync.Mutex
}

type HubManager struct {
	Hubs map[int]*HubInstance // hubID -> HubInstance
	mu   sync.RWMutex
}

type HubMessage struct {
	MsgType HubMsgType
}
