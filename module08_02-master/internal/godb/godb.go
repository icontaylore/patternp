package godb

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"strings"
	"time"
)

type Instance struct {
	Db *pgxpool.Pool
}

func (i *Instance) Start() {
	fmt.Println("Project godb started!")

	//i.addUser(context.Background(), "Dmitri", 25, false)
	i.getAllUsers(context.Background())

	row, err := i.CheckVerif(context.Background())
	if err != nil {
		log.Printf("ошибка в получении столбцов")
	}

	i.GetFams(context.Background(), row, "Swarowsky")
	i.DeleteNotVerif(context.Background())
}

// Структура пользователя
type User struct {
	Name     string
	Age      int
	IsVerify bool
}

func (i *Instance) updateUserAge(ctx context.Context, name string, age int) {
	_, err := i.Db.Exec(ctx, "UPDATE users SET age=$1 WHERE name=$2;", age, name)
	if err != nil {
		fmt.Println(err)
		return
	}
}

func (i *Instance) getUserByName(ctx context.Context, name string) {
	//Выполнение самого запроса. И получение структуры rows, которая содержит в себе строки из базы данных.
	user := &User{}
	err := i.Db.QueryRow(ctx, "SELECT name, age, verify FROM users WHERE name=$1 LIMIT 1;", name).Scan(&user.Name, &user.Age, &user.IsVerify)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("User by name: %v\n", user)
}

// Функция, которая получает список пользователей
func (i *Instance) getAllUsers(ctx context.Context) {
	//Определяем слайс users, куда будем складывать всех пользователей, которых получим из базы
	var users []User

	//Выполнение самого запроса. И получение структуры rows, которая содержит в себе строки из базы данных.
	rows, err := i.Db.Query(ctx, "SELECT name, age, verify FROM users;")
	if err != nil {
		fmt.Println(err)
		return
	}
	//После того как все действия со строками будут выполнены, обязательно и всегда нужно закрывать структуру rows. Для избежания утечек памяти и утечек конектов к базе
	defer rows.Close()

	for rows.Next() {
		user := User{}
		rows.Scan(&user.Name, &user.Age, &user.IsVerify)
		users = append(users, user)
	}

	fmt.Println(users)
}

func (i *Instance) addUser(ctx context.Context, name string, age int, isVerify bool) {
	commandTag, err := i.Db.Exec(ctx, "INSERT INTO users (created_at, name, age, verify) VALUES ($1, $2, $3, $4)", time.Now(), name, age, isVerify)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(commandTag.String())
	fmt.Println(commandTag.RowsAffected())
}

func (i *Instance) CheckVerif(ctx context.Context) (pgx.Rows, error) {
	row, err := i.Db.Query(ctx, "SELECT name FROM users WHERE verify = true")
	if err != nil {
		log.Printf("ошибка в выполнении запроса на вериф")
		return nil, err
	} else if err == pgx.ErrNoRows {
		log.Printf("нет столбцов")
		return nil, err
	}
	return row, nil
}
func (i *Instance) CheckVerifalse(ctx context.Context) (pgx.Rows, error) {
	row, err := i.Db.Query(ctx, "SELECT name FROM users WHERE verify = false")
	if err != nil {
		log.Printf("ошибка в выполнении запроса на вериф")
		return nil, err
	} else if err == pgx.ErrNoRows {
		log.Printf("нет столбцов")
		return nil, err
	}
	return row, nil
}

func (i *Instance) GetFams(ctx context.Context, rows pgx.Rows, fams string) error {
	for rows.Next() {
		var name string
		rows.Scan(&name)
		b := strings.Builder{}

		b.WriteString(name)
		b.WriteString(" ")
		b.WriteString(fams)
		i.Db.Exec(ctx, "UPDATE users SET name=$1 WHERE name=$2", b.String(), name)
		b.Reset()
	}
	defer rows.Close()
	return nil
}

func (i *Instance) DeleteNotVerif(ctx context.Context) {
	row, err := i.CheckVerifalse(ctx)
	if err != nil {
		log.Printf("ошибка")
	}
	defer row.Close()
	for row.Next() {
		var name string
		row.Scan(&name)
		i.Db.Exec(ctx, "DELETE FROM users WHERE name=$1", name)
	}

}
