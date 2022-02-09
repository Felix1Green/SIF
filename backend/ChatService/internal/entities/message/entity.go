package message

type ContentType int
type MsgDirection string

const (
	TextPlain ContentType = iota
	TextMarkdown
	TextHtml
	Image
	Gallery
	Kml
	Attachment
	AttachmentList
	Video
	VCard
	Other = 255
)

const (
	Incoming MsgDirection = "incoming"
	OutGoing MsgDirection = "outgoing"
)

type Message struct {
	ConversationId   int64        `json:"conversationId"`
	Id               int64        `json:"id"`
	SenderId         int64        `json:"senderId"`
	Content          string       `json:"content"`
	ContentType      ContentType  `json:"contentType"`
	MessageDirection MsgDirection `json:"messageDirection"`
}
