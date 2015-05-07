package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pingPeriod = 55 * time.Second
	httpPort   = 8080
	httpHost   = "localhost"
	socketPath = "/socket"
)

type TplData struct {
	MapApiKey string
	SocketUri string
}

func Home(tpl template.Template) func(w http.ResponseWriter, req *http.Request) {
	tplData := TplData{
		MapApiKey: ApiKey,
		SocketUri: fmt.Sprintf("ws://%s:%d%s", httpHost, httpPort, socketPath),
	}
	return func(w http.ResponseWriter, req *http.Request) {
		err := tpl.Execute(w, &tplData)
		if err != nil {
			log.Println(err)
		}
	}
}

func Socket(w http.ResponseWriter, req *http.Request) {
	var upgrader websocket.Upgrader
	pingTicker := time.NewTicker(pingPeriod)
	log.Println("Opening socket")
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}

		return
	}

	ge := GeoEvent{-36.8484597, 174.7633315}
	go ge.Send(ws, &ge)

	time.Sleep(time.Second * 5)

	ge = GeoEvent{59.32522, 18.07002}
	go ge.Send(ws, &ge)

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
		log.Println(err)
	}

	http.HandleFunc(socketPath, Socket)
	http.HandleFunc("/", Home(*page))
	http.ListenAndServe(fmt.Sprintf("%s:%d", httpHost, httpPort), nil)
}
