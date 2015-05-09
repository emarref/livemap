package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

const (
	pingPeriod = 55 * time.Second
	httpPort   = 8080
	httpHost   = "localhost"
	socketPath = "/socket"
)

var socket *websocket.Conn

func main() {
	c := &controller{}
	http.HandleFunc(socketPath, c.socket())
	http.HandleFunc("/", c.home())
	http.ListenAndServe(fmt.Sprintf("%s:%d", httpHost, httpPort), nil)

	go provideEvents()
}
