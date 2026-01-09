package controllers

import (
	"log"

	"github.com/gofiber/websocket/v2"
	"github.com/maonks/absen-rfid-backend/models"
)

var Clients = map[*websocket.Conn]bool{}

func WebsocketHandler(c *websocket.Conn) {
	log.Println("🟢 WS CONNECTED")

	Clients[c] = true
	
	log.Println("TOTAL WS CLIENT:", len(Clients))

	defer func() {
		delete(Clients, c)
		log.Println("🔴 WS DISCONNECTED")
		c.Close()
	}()

	for {
		if _, _, err := c.ReadMessage(); err != nil {
			log.Println("❌ WS READ ERROR:", err)
			break
		}
	}
}

func Broadcast(data models.RealTime) {
	log.Println("📡 BROADCAST TO", len(Clients), "CLIENTS")

	for c := range Clients {
		err := c.WriteJSON(data)
		if err != nil {
			log.Println("❌ WS WRITE ERROR:", err)
			delete(Clients, c)
		}
	}
}
