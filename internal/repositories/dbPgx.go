package repositories

import (
	"context"
	"github.com/jackc/pgx/v4"
	"log"
	"time"
)

type Storage struct {
	conn *pgx.Conn
}

func (d *Storage) DatabaseStart() (*Storage, error) {
	connStr := "postgres://postgres:1@localhost:5432/postgres"
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()

	conn, err := pgx.Connect(ctx, connStr)
	if err != nil {
		log.Printf("ошибка в подкл. к дб: %v", err)
		return nil, err
	}

	if err = conn.Ping(ctx); err != nil {
		log.Printf("нет пинга: %v", err)
		return nil, err
	}

	log.Println("cоединение с БД установлено успешно")

	d.conn = conn
	return d, nil
}
