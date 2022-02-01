package chat_interactor

import (
	"ChatService/internal"
	"ChatService/internal/components/chat_session"
	"ChatService/internal/entities/message"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"sync"
)

type chatInteractor struct {
	usersPeer    map[int64]*websocket.Conn
	userMessages map[int64]map[int64]message.Message
	log          *logrus.Logger
	mu           *sync.Mutex
}

func New(logger *logrus.Logger, mutex *sync.Mutex) *chatInteractor {
	return &chatInteractor{
		usersPeer:    make(map[int64]*websocket.Conn, 0),
		log:          logger,
		userMessages: make(map[int64]map[int64]message.Message, 0),
		mu:           mutex,
	}
}

func (u *chatInteractor) SetNewPeer(userId int64, peer *websocket.Conn) {
	u.usersPeer[userId] = peer

	u.log.Infof("starting new session with user: %d", userId)
	go chat_session.NewChatSession(userId, peer, u.log, u.userMessages[userId], u.eventCallback)
}

func (u *chatInteractor) eventCallback(userId int64) (*websocket.Conn, error) {
	v, ok := u.usersPeer[userId]
	if !ok {
		return nil, internal.UserPeerNotFoundError
	}

	return v, nil
}
