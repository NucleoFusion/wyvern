package mgrctx

import (
	"database/sql"
	"fmt"

	"wyvern-server/internal/managers/channel"
	"wyvern-server/internal/managers/hub"
)

var ctx *MgrCtx

type MgrCtx struct {
	HubMgr   *hub.HubManager
	HubChan  chan *hub.HubMessage
	ChanMgr  *channel.ChannelManager
	ChanChan chan *channel.ChannelMessage
}

func GetCtx() *MgrCtx {
	return ctx
}

func SetCtx(pg *sql.DB) {
	fmt.Println("[Managers] Creating Managers....")
	hubChan := make(chan *hub.HubMessage)
	chanChan := make(chan *channel.ChannelMessage)
	fmt.Println("[Managers] Created Interaction Channels")

	fmt.Println("[Managers] Creating Manager Instances....")
	hm := hub.NewHubManager(pg)
	cm := channel.NewChannelManager(pg)
	fmt.Println("[Managers] Created Manager Instances")

	fmt.Println("[Managers] Running Manager Routines....")
	go hm.Run(hubChan)
	go cm.Run(chanChan)
	fmt.Println("[Managers] Succesfully Ran Manager Routines")

	ctx = &MgrCtx{
		HubMgr:   hm,
		HubChan:  hubChan,
		ChanMgr:  cm,
		ChanChan: chanChan,
	}
}
