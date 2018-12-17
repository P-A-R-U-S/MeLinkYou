package Handlers

import (
	"MeLinkYou/Services/Authentication/Server/Commands"
	"MeLinkYou/Services/Authentication/Server/Data"
	"MeLinkYou/Services/Authentication/Server/Events"
	"cloud.google.com/go/datastore"
	"errors"
	"golang.org/x/net/context"
	"log"
	"reflect"
	"time"
)

type ActivateUserHandler struct {
	name, event string
	Client      *datastore.Client
	Context     context.Context
}

func (h *ActivateUserHandler) Event() string {
	return Events.ACTIVATE_USER_EVENT
}

func (h *ActivateUserHandler) Execute(command Commands.Command) error {

	// validation
	if command == nil {
		return errors.New("no required  arguments passed into handler:" + h.name)
	}

	// get Command
	c, ok := command.(Commands.AddUserCommand)

	if !ok {

		return errors.New(`Argument has incorrect type:` +
			reflect.TypeOf(command).String() +
			`Required instance of ` + h.name)
	}

	user := &Data.User{
		Email:      c.Email,
		Password:   c.Password,
		Created:    time.Now(),
		LastLogged: time.Now(),
	}

	// adds a user with the given email and password to the datastore,
	// returning the key of the newly created entity.
	key := datastore.NameKey("User", user.Email, nil)
	_, err := h.Client.Put(h.Context, key, user)
	if err != nil {
		log.Printf("AddUserHandler.Execute --> Error save user:%s\n", err)
		return err
	}
	log.Printf("AddUserHandler.Execute: Saved User:" + user.Email)

	return nil
}

func (h *ActivateUserHandler) OnSubscribe() {

	h.name = reflect.TypeOf(h).String()

	log.Printf("Handler: %s subscribed.", h.name)
}

func (h *ActivateUserHandler) OnUnsubscribe() {
	h.Client = nil
	h.Context = nil

	log.Printf("Handler: %s unsubscribed.", h.name)
}
