package message

type Message struct{
	Id int64
	Status string
	ContentType string
	SenderId int64
	Direction string
	Content string
}
