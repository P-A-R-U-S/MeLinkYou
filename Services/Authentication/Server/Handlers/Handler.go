package Handlers

import "MeLinkYou/Services/Authentication/Server/Commands"

type Handler interface {
	Event() string
	Execute(command Commands.Command) error
	OnSubscribe()
	OnUnsubscribe()
}
