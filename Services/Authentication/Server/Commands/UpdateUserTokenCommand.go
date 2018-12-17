package Commands

import (
	"MeLinkYou/Services/Authentication/Server/Events"
)

type UpdateUserTokenCommand struct {
	Email string
	Token string
}

func (c UpdateUserTokenCommand) Event() string {
	return Events.UPDATE_USER_TOKEN_EVENT
}
