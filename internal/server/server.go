package server

import (
	"log"
	"net/http"
)

type HttpServ struct {
	server http.Server
}

func GetServer() *HttpServ {
	return &HttpServ{}
}

func (h *HttpServ) Run() {
	var err error
	newServer := http.Server{
		Addr: "localhost:8080",
	}

	log.Println("сервер работает")

	if err = newServer.ListenAndServe(); err != nil {
		log.Printf("сервер не работает")
	}
}
