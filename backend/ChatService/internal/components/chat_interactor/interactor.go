package chat_interactor

import (
	"ChatService/internal"
	"ChatService/internal/components/chat_session"
	"ChatService/internal/entities/chat_companions"
	"ChatService/internal/entities/message"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
)

type chatInteractor struct {
	usersPeer    map[int64]*websocket.Conn
	userMessages map[chat_companions.ChatCompanions][]message.Message
	log          *logrus.Logger
}

func New(logger *logrus.Logger) *chatInteractor {
	return &chatInteractor{
		usersPeer: make(map[int64]*websocket.Conn, 0),
		log:       logger,
	}
}

func (u *chatInteractor) SetNewPeer(userId int64, peer *websocket.Conn) {
	u.usersPeer[userId] = peer

	u.log.Infof("starting new session with user: %d", userId)
	go chat_session.NewChatSession(userId, peer, u.log, u.userMessages[chat_companions.ChatCompanions{
		User1: userId,
	}], u.eventCallback)
}

func (u *chatInteractor) eventCallback(userId int64) (*websocket.Conn, error) {
	v, ok := u.usersPeer[userId]
	if !ok {
		return nil, internal.UserPeerNotFoundError
	}

	return v, nil
}
