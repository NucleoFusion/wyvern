package hub

import (
	"database/sql"
	"sync"

	"wyvern-server/internal/managers/channel"
	"wyvern-server/internal/models"
)

type HubMsgType int

const (
	AddHub HubMsgType = iota
	RemoveHub
	AddChannel
	RemoveChannel
	RegisterClient
	UnregisterClient
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
	Pg   *sql.DB
}

type HubMessage struct {
	MsgType HubMsgType
	Param   any
}

type HubClient struct {
	HubID    int
	ClientID int
}
