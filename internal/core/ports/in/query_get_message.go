package in

import tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"

type QueryMessageUpdateHandler interface {
	HandleUpdates(updates tgbotapi.UpdatesChannel)
}

type MessageUpdateHandler interface {
	HandleUpdates(updates tgbotapi.UpdatesChannel)
}
