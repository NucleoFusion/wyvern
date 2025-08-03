package handlers

import (
	"encoding/json"
	"fmt"

	"wyvern-server/internal/dbctx/mgrctx"
	"wyvern-server/internal/log"
	"wyvern-server/internal/managers/hub"
	"wyvern-server/internal/models"

	"github.com/gorilla/websocket"
)

// When are these goroutines stopped?
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

func HandleClientWrites(m *models.Client) {
	for {
		select {
		case msg, ok := <-m.Send:
			if !ok {
				_ = m.Conn.WriteMessage(websocket.CloseMessage, []byte{})
				return
			}

			err := m.Conn.WriteMessage(websocket.TextMessage, msg)
			if err != nil {
				log.Log(log.Moderate, fmt.Sprintf("Could not send message to %d: %s", m.UserID, err.Error()))
			}
		}
	}
}
