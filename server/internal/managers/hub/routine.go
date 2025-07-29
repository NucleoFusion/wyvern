package hub

import (
	"wyvern-server/internal/log"
	"wyvern-server/internal/models"
)

func (m *HubManager) Run(msgChan chan *HubMessage) {
	for msg := range msgChan {
		switch msg.MsgType {
		case AddHub:
			v := msg.Param
			prm, ok := v.(*models.Hub)
			if !ok {
				log.Log(log.Moderate, "Invalid Type for Params - AddHub (HubManager)")
				continue
			}

			go m.AddHub(prm)
		case RemoveHub:
			v := msg.Param
			prm, ok := v.(*models.Hub)
			if !ok {
				log.Log(log.Moderate, "Invalid Type for Params - RemoveHub (HubManager)")
				continue
			}

			go m.RemoveHub(prm)
		case AddChannel:
			v := msg.Param
			prm, ok := v.(*models.Channel)
			if !ok {
				log.Log(log.Moderate, "Invalid Type for Params - AddChannel (HubManager)")
				continue
			}

			go m.AddChannel(prm)
		case RemoveChannel:
			v := msg.Param
			prm, ok := v.(*models.Channel)
			if !ok {
				log.Log(log.Moderate, "Invalid Type for Params - RemoveChannel (HubManager)")
				continue
			}

			go m.RemoveChannel(prm)
		}
	}
}
