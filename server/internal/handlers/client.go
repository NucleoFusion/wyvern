package handlers

import (
	"encoding/json"

	"wyvern-server/internal/dbctx/mgrctx"
	"wyvern-server/internal/log"
	"wyvern-server/internal/managers/hub"
	"wyvern-server/internal/models"
)

func HandleClientReads(m *models.Client) {
	defer func() {
		m.Conn.Close()
	}()

	for {
		_, msg, err := m.Conn.ReadMessage()
		if err != nil {
			log.Log(log.Moderate, err.Error())
			continue
		}

		var m models.Message
		if err := json.Unmarshal(msg, &m); err != nil {
			log.Log(log.Moderate, "Invalid Message Format")
			continue
		}

		mgrctx.GetCtx().HubChan <- &hub.HubMessage{
			MsgType: hub.IncomingMsg,
			Param:   &m,
		}

	}
}
