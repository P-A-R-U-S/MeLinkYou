package Commands

import "MeLinkYou/Services/Authentication/Server/Events"

type AddUserCommand struct {
	Email    string
	Password string
}

func (c AddUserCommand) Event() string {
	return Events.ADD_USER_EVENT
}
