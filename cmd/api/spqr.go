package main

import (
	"galileoio_bot/internal/adapters/in/consumer"
	"galileoio_bot/internal/adapters/out/producer"
	"galileoio_bot/internal/adapters/out/telegram"
	"galileoio_bot/internal/core/usecases"
)

func main() {
	updates := telegram.InitBot()

	messageSender := producer.TelegramMessageSender{}
	sendMessageCommand := usecases.NewSendMessageUseCase(messageSender)

	messageConsumer := consumer.NewMessageConsumer(sendMessageCommand)
	messageConsumer.HandleUpdates(updates)
}
