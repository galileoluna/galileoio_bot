package domain

type Message struct {
	ChatID           int64
	Text             string
	ReplyToMessageID int
}
