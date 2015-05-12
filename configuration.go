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
	Map struct {
		ApiKey string `yaml:"api_key"`
	}
}

var cfg Config

func AppendConfiguration(filename string) error {
	log.Printf("Loading configuration file \"%s\"", filename)
	content, err := ioutil.ReadFile(filename)

	if err != nil {
		// If a file is not found, log and skip
		log.Println(err)
		return nil
	}

	log.Println("Unmarshaling", len(content), "bytes of configuration")
	err = yaml.Unmarshal(content, &cfg)

	if err != nil {
		return err
	}

	return nil
}

func InitialiseConfiguration(filenames ...string) error {
	for _, filename := range filenames {
		err := AppendConfiguration(filename)

		if err != nil {
			return err
		}
	}

	return nil
}
