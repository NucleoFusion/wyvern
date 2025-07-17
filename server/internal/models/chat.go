package models

type Hub struct {
	ID      int    `json:"id" db:"id"`
	Repo    string `json:"repo" db:"repo"`
	OwnerID int    `json:"owner_id" db:"owner_id"`
}

type Channel struct {
	ID      int    `json:"id" db:"id"`
	HubID   int    `json:"hub_id" db:"hub_id"` // references Hub.ID
	Name    string `json:"name" db:"name"`
	Private bool   `json:"private" db:"private"`
}

type Message struct {
	ID        int    `json:"id"`
	HubID     int    `json:"hub_id"`
	UserID    int    `json:"user_id"`
	Content   string `json:"content"`
	Timestamp int64  `json:"timestamp"`
}
