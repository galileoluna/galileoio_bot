package producer

import (
	"galileoio_bot/internal/adapters/out/telegram"
	"galileoio_bot/internal/core/domain"
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

type TelegramMessageSender struct{}

func (t TelegramMessageSender) SendMessage(message domain.Message) error {
	msg := tgbotapi.NewMessage(message.ChatID, message.Text)
	msg.ReplyToMessageID = message.ReplyToMessageID

	_, err := telegram.Bot.Send(msg)
	return err
}
