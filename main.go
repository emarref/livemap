package main

import "log"
import "time"

type GeoEvent struct {
	Lat, Long float32
}

var socket chan GeoEvent

func main() {
	// Acquire configuration
	log.Println("Initialising configuration")
	err := InitialiseConfiguration("config.yml")
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Open websocket
	log.Println("Initialising websocket")
	wc, err := InitialiseWebsocket()
	if err != nil {
		log.Fatalln(err)
		return
	}

	// Initialise provider for generating geo events
	// First provider will be sending dummy events every 3 seconds.
	log.Println("Initialising provider")
	provider := InitialiseProvider(wc)
	go func() {
		poll := time.NewTicker(time.Second * 5)
		for {
			select {
			case <-poll.C:
				provider.Tick()
			case msg := <-wc.In:
				wc.Out <- msg
			}
		}
	}()

	// Serve home page
	log.Println("Initialising web server")
	err = InitialiseWebserver()
	if err != nil {
		log.Fatalln(err)
		return
	}
}
