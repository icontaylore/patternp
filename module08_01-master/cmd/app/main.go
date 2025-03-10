package main

import (
	"context"
	"fmt"
	"net/url"
	"os"

	//Импортируем пакет для работы с пулом соединений
	"github.com/jackc/pgx/v4/pgxpool"
)

func main() {
	conStr :=
		fmt.Sprintf("%s://%s:%s@%s:%s/%s", "postgres",
			url.QueryEscape("db_user"),
			url.PathEscape("1"),
			"localhost",
			"54320",
			"db_test",
		)

	ctx, _ := context.WithCancel(context.Background())

	//Сконфигурируем пул, задав для него максимальное количество соединений
	poolConfig, _ := pgxpool.ParseConfig(conStr)
	poolConfig.MaxConns = 5

	//Получаем пул соединений, используя контекст и конфиг
	pool, err := pgxpool.ConnectConfig(ctx, poolConfig)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Connect to database failed: %v\n", err)
		os.Exit(1)
	}
	fmt.Println("Connection OK!")

	for i := 0; i < 5; i++ {
		go func(count int) {
			//Функция Exec выполняет запрос к БД.
			_, err = pool.Exec(ctx, ";")
			if err != nil {
				fmt.Fprintf(os.Stderr, "Ping failed: %v\n", err)
				os.Exit(1)
			}
			fmt.Println(count, "Query OK!")

			//Выводим информацию о соединениях
			fmt.Printf("Conections - Max: %d, Iddle: %d, Total: %d \n", pool.Stat().MaxConns(), pool.Stat().IdleConns(), pool.Stat().TotalConns())
		}(i)
	}

	select {}

}
