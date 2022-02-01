package chat_session

import (
	"ChatService/internal"
	"ChatService/internal/entities/message"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type ChatSession struct {
	userID   int64
	peer     *websocket.Conn
	log      *logrus.Logger
	callback internal.EventCallback
}

func NewChatSession(
	userID int64,
	peer *websocket.Conn,
	log *logrus.Logger,
	messageHistory []message.Message,
	callback internal.EventCallback,
) *ChatSession {
	return &ChatSession{
		userID:   userID,
		peer:     peer,
		log:      log,
		callback: callback,
	}
}

func (s *ChatSession) Start() {

}
