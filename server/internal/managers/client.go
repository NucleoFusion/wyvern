package managers

import (
	"fmt"

	"wyvern-server/internal/log"
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

		// var m models.Message
		// if err := json.Unmarshal(msg, &m); err != nil {
		// 	log.Log(log.Moderate, "Invalid Message Format")
		// 	continue
		// }

		fmt.Println(string(msg))
	}
}
