package models

type HubRoutine struct {
	Add        chan *Hub
	RemoveHub  chan *int
	RemoveUser chan *int
}
