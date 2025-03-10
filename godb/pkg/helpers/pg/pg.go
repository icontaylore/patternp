package pg

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v4/pgxpool"
	"log"
	"net/url"
)

type Config struct {
	Host     string
	Port     string
	Username string
	Password string
	DbName   string
	Timeout  int
}

func InitConfig(c *Config) (*pgxpool.Config, error) {
	connStr := fmt.Sprintf("%s://%s:%s@%s:%s/%s", "postgres",
		url.QueryEscape(c.Username),
		url.QueryEscape(c.Password),
		c.Host,
		c.Port,
		c.DbName,
	)

	conf, err := pgxpool.ParseConfig(connStr)
	if err != nil {
		log.Printf("ошибок нет")
		return nil, err
	}

	return conf, nil
}

func CreatePoolConnect(poolConf *pgxpool.Config) (*pgxpool.Pool, error) {
	conn, err := pgxpool.ConnectConfig(context.Background(), poolConf)
	if err != nil {
		log.Printf("ошибка в подключении")
		return nil, err
	}
	return conn, nil
}
