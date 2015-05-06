package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/emarref/livemap/mapping"
	"github.com/gorilla/websocket"
)

const (
	writeWait  = 10 * time.Second
	pingPeriod = 55 * time.Second
	httpPort   = "8080"
	httpHost   = "127.0.0.1"
	socketPath = "/socket"
)

type TplData struct {
	MapApiKey string
	SocketUri string
}

func Home(tpl template.Template) func(w http.ResponseWriter, req *http.Request) {
	tplData := TplData{
		MapApiKey: mapping.ApiKey,
		SocketUri: fmt.Sprintf("ws://%s:%s%s", httpHost, httpPort, socketPath),
	}
	return func(w http.ResponseWriter, req *http.Request) {
		err := tpl.Execute(w, &tplData)
		if err != nil {
			log.Println(err)
		}
	}
}

func Socket(w http.ResponseWriter, req *http.Request) {
	log.Println("Opening socket")
	pingTicker := time.NewTicker(pingPeriod)
	upgrader := websocket.Upgrader{ReadBufferSize: 1024, WriteBufferSize: 1024}
	ws, err := upgrader.Upgrade(w, req, nil)
	if err != nil {
		if _, ok := err.(websocket.HandshakeError); !ok {
			log.Println(err)
		}

		return
	}

	ge := mapping.GeoEvent{-36.8484597, 174.7633315}
	go ge.Send(ws, &ge)

	time.Sleep(time.Second * 5)

	ge = mapping.GeoEvent{59.32522, 18.07002}
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
	http.ListenAndServe(fmt.Sprintf("%s:%s", httpHost, httpPort), nil)
}
