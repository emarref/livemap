package main

import (
	"log"

	"github.com/gorilla/websocket"
)

const (
	ApiKey = "AIzaSyBID81MgktFhz0F4T_Klb7YlwvxSvJrqPY"
)

type GeoEvent struct {
	Lat, Long float64
}

func (ge *GeoEvent) Send(ws *websocket.Conn, geoEvent *GeoEvent) {
	log.Println("Sending geo event", geoEvent)

	if err := ws.WriteJSON(geoEvent); err != nil {
		log.Println("Could not send message")
		log.Println(err)
	}
}
