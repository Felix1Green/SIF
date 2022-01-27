package chat_session

import (
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type ChatSession struct{
	userID int64
	peer *websocket.Conn
	log *logrus.Logger
}

func NewChatSession(userID int64, peer *websocket.Conn, log *logrus.Logger) *ChatSession{
	return &ChatSession{
		userID: userID,
		peer: peer,
		log: log,
	}
}

func (s *ChatSession) Start(){

}
