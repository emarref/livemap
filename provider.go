package main

import (
	"encoding/json"
	"log"
)
import "github.com/emarref/webchannel"
import "math/rand"
import "time"

type Provider interface {
	Tick()
}

type DummyProvider struct {
	wc *webchannel.WebChannel
}

func (provider DummyProvider) Tick() {
	rand.Seed(time.Now().UnixNano())
	geo := GeoEvent{
		Lat:  rand.Float32()*360 - 180,
		Long: rand.Float32()*360 - 180,
	}

	geoJson, err := json.Marshal(geo)

	if err != nil {
		log.Fatalln(err)
		return
	}

	provider.wc.Out <- geoJson
}

func InitialiseProvider(wc *webchannel.WebChannel) Provider {
	return DummyProvider{wc}
}
