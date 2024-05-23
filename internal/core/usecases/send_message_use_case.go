package usecases

import (
	"galileoio_bot/internal/core/domain"
	"galileoio_bot/internal/core/ports/in"
)

type SendMessageUseCase struct {
	messageSender in.CommandMessageSender
}

func NewSendMessageUseCase(messageSender in.CommandMessageSender) in.CommandSendMessage {
	return &SendMessageUseCase{messageSender: messageSender}
}

func (u *SendMessageUseCase) Handle(message domain.Message) error {
	return u.messageSender.SendMessage(message)
}
