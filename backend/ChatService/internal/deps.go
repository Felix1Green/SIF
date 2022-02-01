package internal

import (
	"fmt"
	"github.com/gorilla/websocket"
)

type EventCallback func(int64) (*websocket.Conn, error)

var (
	UserPeerNotFoundError = fmt.Errorf("user peer not found")
)
