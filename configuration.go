package main

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Config struct {
	Http struct {
		Port int
		Host string
	}
	Socket struct {
		Path string
	}
}

var cfg Config

func InitialiseConfiguration(filename string) error {
	log.Printf("Loading configuration file \"%s\"", filename)
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		return err
	}

	log.Println("Unmarshaling", len(content), "bytes of configuration")
	err = yaml.Unmarshal(content, &cfg)

	if err != nil {
		return err
	}

	return nil
}
