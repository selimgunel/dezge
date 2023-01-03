package couch_test

import (
	"context"
	"testing"

	"github.com/narslan/dezge/couch"
)

func TestNewDB(t *testing.T) {
	client, err := couch.NewClient()
	if err != nil {
		t.Fatal(err)
	}
	_, err = couch.NewDB(context.TODO(), client)
	if err != nil {
		t.Fatal(err)
	}

}
