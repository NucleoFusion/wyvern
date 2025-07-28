package channel

import (
	"sync"

	"wyvern-server/internal/models"
)

type ChannelMsgType int

const (
	AddChannel ChannelMsgType = iota
	RemoveChannel
	AddClient
	RemoveClient
)

type ChannelInstance struct {
	ID      int
	Clients map[*models.Client]bool // Map for efficient lookups
	mu      sync.Mutex
}

type ChannelManager struct {
	Channels map[int]*ChannelInstance // ChannelID to instance
	mu       sync.RWMutex
}

type ChannelMessage struct {
	MsgType ChannelMsgType
	Param   any
}
