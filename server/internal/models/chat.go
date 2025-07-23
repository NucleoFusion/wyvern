package models

import (
	"time"

	"github.com/gorilla/websocket"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

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
	ID        primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	ChannelID int                `json:"channel_id" bson:"channel_id"`
	UserID    int                `json:"user_id" bson:"user_id"`
	Content   string             `json:"content" bson:"content"`
	Timestamp time.Time          `json:"timestamp" bson:"timestamp"`
}

type Client struct {
	Conn      *websocket.Conn
	Send      chan []byte
	UserID    int
	ChannelID int // References ChannelID
}
