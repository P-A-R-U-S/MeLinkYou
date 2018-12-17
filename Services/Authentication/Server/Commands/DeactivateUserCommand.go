package Commands

import "MeLinkYou/Services/Authentication/Server/Events"

type DeactivateUserCommand struct {
	Email string
	Token string
}

func (c DeactivateUserCommand) Event() string {
	return Events.DEACTIVATE_USER_EVENT
}
