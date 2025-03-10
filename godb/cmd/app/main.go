package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"servise/godb/internal/godb"
	"servise/godb/pkg/helpers/pg"
)

func main() {
	conf := &pg.Config{
		DbName:   "db_test",
		Username: "db_user",
		Host:     "localhost",
		Port:     "54320",
		Password: "1",
	}
	pgxConfig, err := pg.InitConfig(conf)
	if err != nil {
		log.Printf("ошибка в получении конфига")
		os.Exit(1)
	}

	// установим максимальное кол-во соед.
	pgxConfig.MaxConns = 5

	pgxPoolDb, err := pg.CreatePoolConnect(pgxConfig)
	if err != nil {
		log.Printf("ошибка в создании пула")
		os.Exit(1)
	}
	fmt.Printf("конект есть\n")

	// чекаем на соединение
	if _, err = pgxPoolDb.Exec(context.Background(), ";"); err != nil { // ; служит пинговой штукой
		log.Printf("ошибка в пинге")
		os.Exit(1)
	}

	ins := &godb.Instance{DB: pgxPoolDb}
	ins.Start()
}
