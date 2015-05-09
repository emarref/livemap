package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type TplData struct {
	MapApiKey string
	SocketUri string
}

type controller struct {
}

func (c *controller) home() func(http.ResponseWriter, *http.Request) {
	tpl, err := template.ParseFiles("templates/index.html")

	if err != nil {
		log.Fatalln(err)
	}

	data := TplData{
		MapApiKey: apiKey,
		SocketUri: fmt.Sprintf("ws://%s:%d%s", httpHost, httpPort, socketPath),
	}

	return func(w http.ResponseWriter, req *http.Request) {
		tpl.Execute(w, data)
	}
}

func (c *controller) socket() func(http.ResponseWriter, *http.Request) {
	var upgrader websocket.Upgrader
	ping := time.NewTicker(pingPeriod)

	return func(w http.ResponseWriter, req *http.Request) {
		log.Println("Opening socket")

		var err error
		socket, err = upgrader.Upgrade(w, req, nil)

		if nil != err {
			log.Fatalln(err)
			return
		}

		log.Println("Socket open")

		defer func() {
			log.Println("Closing socket")
			socket.Close()
			ping.Stop()
		}()

		go func() {
			for t := range ping.C {
				log.Println(t)
				socket.SetWriteDeadline(time.Now().Add(writeWait))
				if err := socket.WriteMessage(websocket.PingMessage, []byte{}); err != nil {
					return
				}
			}
			log.Println("Left loop")
		}()
	}
}
