package telegram

import (
	"log"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

var Bot *tgbotapi.BotAPI

func InitBot() tgbotapi.UpdatesChannel {
	var err error
	Bot, err = tgbotapi.NewBotAPI("6371396654:AAEY2jo3-1auxpxpW609UGWplmk9-CipK0g")
	if err != nil {
		log.Panic(err)
	}

	Bot.Debug = true

	updateConfig := tgbotapi.NewUpdate(0)
	updateConfig.Timeout = 30

	updates := Bot.GetUpdatesChan(updateConfig)
	return updates
}
