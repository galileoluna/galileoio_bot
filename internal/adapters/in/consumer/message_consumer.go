package consumer

import (
	"galileoio_bot/internal/core/domain"
	"galileoio_bot/internal/core/ports/in"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"log"
)

type MessageConsumer struct {
	sendMessageCommand in.CommandSendMessage
}

func NewMessageConsumer(sendMessageCommand in.CommandSendMessage) in.MessageUpdateHandler {
	return &MessageConsumer{sendMessageCommand: sendMessageCommand}
}

func (c *MessageConsumer) HandleUpdates(updates tgbotapi.UpdatesChannel) {
	for update := range updates {
		if update.Message == nil {
			continue
		}

		message := domain.Message{
			ChatID:           update.Message.Chat.ID,
			Text:             update.Message.Text,
			ReplyToMessageID: update.Message.MessageID,
		}

		err := c.sendMessageCommand.Handle(message)
		if err != nil {
			log.Println(err)
		}
	}
}
