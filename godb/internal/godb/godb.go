package godb

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"time"
)

type Instance struct {
	DB *pgxpool.Pool
}

type User struct {
	name   string
	age    int
	verify bool
}

func (i *Instance) Start() {
	i.AddUser(context.Background(), "James", 20, false)
	i.GetAllUsers()
}

func (i *Instance) AddUser(ctx context.Context, name string, age int, verify bool) {
	command, err := i.DB.Exec(ctx,
		"INSERT INTO users (created_at,name,age,verify) VALUES ($1,$2,$3,$4)", time.Now(), name, age, verify)
	if err != nil {
		log.Printf("ошибка вставки")
	}
	fmt.Println(command.String())
}

func (i *Instance) GetAllUsers() {
	u := make([]User, 0, 100)
	rows, err := i.DB.Query(context.Background(), "SELECT name,age,verify FROM users")
	if err == pgx.ErrNoRows {
		log.Printf("нету столбцов")
		return
	} else if err != nil {
		log.Printf("ошибка связанная с получением информации из запроса")
	}
	defer rows.Close()

	for rows.Next() {
		user := User{}
		rows.Scan(&user.name, &user.age, &user.verify)
		u = append(u, user)
	}

	fmt.Println(u)
}
