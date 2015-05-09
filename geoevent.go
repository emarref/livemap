package main

import (
	"log"

	"github.com/gorilla/websocket"
)

const (
	apiKey = "AIzaSyBID81MgktFhz0F4T_Klb7YlwvxSvJrqPY"
)

type GeoEvent struct {
	Lat, Long float64
}

func (ge *GeoEvent) Send(ws *websocket.Conn) {
	log.Println("Sending geo event", ge)

	if err := ws.WriteJSON(ge); err != nil {
		log.Println("Could not send message")
		log.Println(err)
	}
}
