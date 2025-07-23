package hub

func (m *HubManager) Run(msgChan *chan HubMessage) {
	for msg := range *msgChan {
		switch msg.MsgType {
		}
	}
}
