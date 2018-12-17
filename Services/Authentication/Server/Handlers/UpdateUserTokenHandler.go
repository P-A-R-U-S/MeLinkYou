package Handlers

import (
	"MeLinkYou/Services/Authentication/Server/Commands"
	"MeLinkYou/Services/Authentication/Server/Events"
	"cloud.google.com/go/datastore"
	"errors"
	"golang.org/x/net/context"
	"log"
	"reflect"
)

type UpdateUserTokenHandler struct {
	name, event, projectID string
	Client                 *datastore.Client
	Context                context.Context
}

func (h *UpdateUserTokenHandler) Event() string {
	return Events.UPDATE_USER_TOKEN_EVENT
}

func (h *UpdateUserTokenHandler) Execute(command Commands.Command) error {

	// validation
	if command == nil {
		return errors.New("no required arguments passed into handler:" + h.name)
	}

	// get Command
	c, ok := command.(Commands.UpdateUserTokenCommand)

	if !ok {

		return errors.New(`Argument has incorrect type:` +
			reflect.TypeOf(c).String() +
			`Required instance of Commands.UpdateUserTokenCommand.`)
	}

	//user := &Data.User{
	//	Email: command.Email,
	//	Password:command.Password }
	//
	//// adds a user with the given email and password to the datastore,
	//// returning the key of the newly created entity.
	//key := datastore.IncompleteKey("User", nil)
	//_, err := h.client.Put(h.context, key, user)
	//if err != nil {
	//	return err
	//}

	return nil
}

func (h *UpdateUserTokenHandler) OnSubscribe() {

	h.name = reflect.TypeOf(h).String()

	log.Printf("Handler: %s subscribed.", h.name)
}

func (h *UpdateUserTokenHandler) OnUnsubscribe() {

	log.Printf("Handler: %s unsubscribed.", h.name)

}
