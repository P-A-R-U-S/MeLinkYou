package Data

import (
	"time"
)

// User is the model used to store userInfo into the datastore.
type User struct {
	Email      string    `datastore:"email"`
	Password   string    `datastore:"password"`
	Created    time.Time `datastore:"created"`
	LastLogged time.Time `datastore:"lastLogged"`
	Token      string    `datastore:"token"`
	IsActive   bool      `datastore:"isActive"`
	Id         string    //The string ID used in the datastore.
}
