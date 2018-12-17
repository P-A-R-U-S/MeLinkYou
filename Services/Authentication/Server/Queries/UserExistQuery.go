package Queries

import (
	"cloud.google.com/go/datastore"
	"golang.org/x/net/context"
)

type UserExistQuery struct {
	projectID string
	Client    *datastore.Client
	Context   context.Context
}

func (q *UserExistQuery) Query(email string) (bool, error) {

	// Create a query to fetch all Users entities, ordered by "created".
	key := datastore.NameKey("User", email, nil)
	query := datastore.NewQuery("User").Filter("__key__ =", key).KeysOnly()

	us, err := q.Client.Count(q.Context, query)

	if err != nil {
		return false, err
	}

	return us > 0, nil
}
