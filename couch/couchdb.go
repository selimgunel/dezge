package couch

import (
	"context"

	_ "github.com/go-kivik/couchdb/v3" // The CouchDB driver
	kivik "github.com/go-kivik/kivik/v3"
)

// GamesDB represents the games documents in couchdb store.
type GamesDB struct {
	*kivik.DB
}

func NewGamesDB(ctx context.Context, client *kivik.Client) (*GamesDB, error) {

	db := client.DB(context.TODO(), "games")

	return &GamesDB{db}, nil
}

// That will be used within a manager.
func NewClient() (*kivik.Client, error) {
	client, err := kivik.New("couch", "http://localhost:5984/")
	if err != nil {
		return nil, err
	}
	return client, nil
}
