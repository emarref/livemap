package main

import (
	"fmt"
	"log"
	"net/http"
)

func InitialiseWebserver() error {
	addr := fmt.Sprintf("%s:%d", cfg.Http.Host, cfg.Http.Port)

	log.Println(fmt.Sprintf("Handling traffic on %s", addr))

	http.Handle("/", http.FileServer(http.Dir("public")))
	return http.ListenAndServe(fmt.Sprintf("%s:%d", cfg.Http.Host, cfg.Http.Port), nil)
}
