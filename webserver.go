package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type TemplateData struct {
	MapApiKey string
	SocketUri string
}

func getAddress() string {
	return fmt.Sprintf("%s:%d", cfg.Http.Host, cfg.Http.Port)
}

func handleHomePage() func(response http.ResponseWriter, request *http.Request) {
	tpl, err := template.ParseFiles("views/index.html")

	if err != nil {
		log.Fatalln(err)
		return nil
	}

	tplData := TemplateData{
		cfg.Map.ApiKey,
		fmt.Sprintf("ws://%s%s", getAddress(), cfg.Socket.Path),
	}

	return func(response http.ResponseWriter, request *http.Request) {
		tpl.Execute(response, tplData)
	}
}

func InitialiseWebserver() error {
	addr := getAddress()

	log.Println(fmt.Sprintf("Handling traffic on %s", addr))

	http.HandleFunc("/", handleHomePage())
	//http.Handle("/", http.FileServer(http.Dir("public")))
	return http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Http.Host, cfg.Http.Port), nil)
}
