package models

import "time"

type Message struct {
	ID     int64     `json:"id_sms"`
	UserId int       `json:"user_id"`
	Cont   string    `json:"cont"`
	time   time.Time `json:"time"`
}

type User struct {
	Id   int64  `json:"id"`
	Name string `json:"name"`
}
