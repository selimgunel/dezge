package sqlite_test

import (
	"context"
	"errors"
	"fmt"
	"os"
	"testing"

	"github.com/narslan/dezge"
	"github.com/narslan/dezge/sqlite"
)

// Te checks if add engine information to the database.
func TestCreateEngineInfo(t *testing.T) {

	db := MustOpenDB(t)
	defer MustCloseDB(t, db)
	dbi := sqlite.NewEngineInfo(db)

	t.Run("ErrPathShouldBePresent", func(t *testing.T) {
		engine := dezge.NewEngineInfo("path-not-exists")
		err := dbi.Create(context.Background(), engine)
		if err != nil {
			var pathError *os.PathError
			ok := errors.As(err, &pathError)
			if !ok {
				t.Fatalf("expected: %T got: %T", pathError, err)
			}
		}
	})

	t.Run("ErrEngineShouldNotBeStoredTwice", func(t *testing.T) {
		engine := dezge.NewEngineInfo(enginePath(t))
		err := dbi.Create(context.Background(), engine)
		assert(err, t)
		err = dbi.Create(context.Background(), engine)
		if err != nil {
			if dezge.ErrorCode(err) != dezge.EINVALID {
				t.Fatalf("expected: %s got: %s", dezge.EINVALID, err.Error())
			}
		}
	})

}

// enginePath returns the path of the engine's executable.
func enginePath(tb testing.TB) string {
	enginePath, ok := os.LookupEnv("UCI_CHESS_ENGINE")
	if !ok {
		err := fmt.Errorf("UCI_CHESS_ENGINE should point to an uci chess engine.")
		assert(err, tb)
	}
	return enginePath
}

// CreateEngineService checks if add engine information to the database.
func TestFindByID(t *testing.T) {

	db := MustOpenDB(t)
	defer MustCloseDB(t, db)
	service := sqlite.NewEngineInfo(db)

	// Save the engines in the database.

	engine := dezge.NewEngineInfo(enginePath(t))
	ctx := context.Background()
	err := service.Create(ctx, engine)
	assert(err, t)

	allEngines, _, err := service.Find(context.Background(), dezge.EngineFilter{})
	assert(err, t)

	expected := allEngines[0].ID
	ctx = context.Background()
	got, err := service.FindByID(ctx, expected)
	assert(err, t)
	if got.ID != expected {
		err = fmt.Errorf("got: %d wanted: %d", got.ID, expected)
		assert(err, t)
	}

}

func assert(err error, t testing.TB) {
	if err != nil {
		t.Fatal(err)
	}
}
