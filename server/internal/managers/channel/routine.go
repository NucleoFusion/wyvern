package channel

import (
	"wyvern-server/internal/log"
	"wyvern-server/internal/models"
)

func (m *ChannelManager) Run(msgChan *chan ChannelMessage) {
	for msg := range *msgChan {
		switch msg.MsgType {
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
