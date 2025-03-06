package main

import (
	"chat-app/internal/repositories"
	"chat-app/internal/server"
	"log"
)

func main() {
	pgx := &repositories.Storage{}
	bd, err := pgx.DatabaseStart()
	if err != nil {
		log.Printf("ошибка в запуске бд", bd)
	}

	s := server.GetServer()
	s.Run()
}
