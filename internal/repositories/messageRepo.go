package repositories

import (
	"chat-app/internal/models"
)

type Messager interface {
	GetMessages()                      // Извлечение сообщений
	AddMessage(message models.Message) // Добавление сообщений
}
