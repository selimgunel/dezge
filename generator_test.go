package dezge_test

import (
	"testing"

	dezge "github.com/narslan/dezge"
)

func TestCounter(t *testing.T) {

	g := dezge.NewGenerator(1)

	if g.Counter() != 2 {
		t.Fatal("expected 2")
	}

}
