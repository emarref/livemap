package main

import "github.com/emarref/webchannel"

func InitialiseWebsocket() (*webchannel.WebChannel, error) {
	wc, err := webchannel.New(cfg.Socket.Path)

	if err != nil {
		return nil, err
	}

	return wc, nil
}
