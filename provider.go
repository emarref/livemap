package main

import (
	"log"
	"time"
)

func provideEvents() {
	timer := time.NewTicker(time.Second * 3)
	go func() {
		for t := range timer.C {
			log.Println(t)
		}
	}()
}
