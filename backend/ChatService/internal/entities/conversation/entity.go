package conversation

import (
	"ChatService/internal/entities/message"
	"ChatService/internal/entities/participant"
)

type Conversation struct {
	ConversationId int64
	Messages       []message.Message
	Description    string
	Participants   []participant.Participant
}
