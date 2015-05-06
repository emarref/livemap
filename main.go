package main

import (
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pongWait   = 60 * time.Second
	pingPeriod = (pongWait * 9) / 10
	filePeriod = 10 * time.Second
)

type GeoEvent struct {
	Lat  float64
	Long float64
}

func Home(tpl template.Template) func(w http.ResponseWriter, req *http.Request) {
	return func(w http.ResponseWriter, req *http.Request) {
		tpl.Execute(w, "Zippity")
	}
}

func SendGeoEvent(ws *websocket.Conn, geoEvent *GeoEvent) {
	log.Println(geoEvent)
	if err := ws.WriteJSON(geoEvent); err != nil {
		log.Println("Could not send message")
		log.Println(err)
	}
}

func Socket(w http.ResponseWriter, req *http.Request) {
	pingTicker := time.NewTicker(pingPeriod)
	upgrader := websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}

		return
	}

	go SendGeoEvent(ws, &GeoEvent{-36.8484597, 174.7633315})
	time.Sleep(time.Second * 5)
	go SendGeoEvent(ws, &GeoEvent{59.32522, 18.07002})
	defer func() {
		log.Println("Closing socket")
		ws.Close()
		pingTicker.Stop()
	}()

	for {
		select {
		case <-pingTicker.C:
			ws.SetWriteDeadline(time.Now().Add(writeWait))
			if err := ws.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
				return
			}
		}
	}
}

func main() {
	page, err := template.ParseFiles("templates/index.html")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/socket", Socket)
	http.HandleFunc("/", Home(*page))
	http.ListenAndServe(":8080", nil)
}
