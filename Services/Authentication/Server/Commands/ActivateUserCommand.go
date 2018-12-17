package Commands

import "MeLinkYou/Services/Authentication/Server/Events"

type ActivateUserCommand struct {
	Email string
	Token string
}

func (c ActivateUserCommand) Event() string {
	return Events.ACTIVATE_USER_EVENT
}
