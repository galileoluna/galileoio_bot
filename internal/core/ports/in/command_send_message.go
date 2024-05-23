package in

import "galileoio_bot/internal/core/domain"

type CommandSendMessage interface {
	Handle(message domain.Message) error
}
