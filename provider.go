package main

import "github.com/emarref/webchannel"

type Provider interface {
	Tick()
}

type DummyProvider struct {
	wc *webchannel.WebChannel
}

func (provider DummyProvider) Tick() {
	provider.wc.Out <- []byte("Dummy provider")
}

func InitialiseProvider(wc *webchannel.WebChannel) Provider {
	return DummyProvider{wc}
}
