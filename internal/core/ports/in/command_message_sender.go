package in

import "galileoio_bot/internal/core/domain"

type CommandMessageSender interface {
	SendMessage(message domain.Message) error
}
