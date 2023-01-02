package dezge

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/google/uuid"
	"github.com/rs/zerolog/log"
)

type GeneratorService interface {
	Counter() uint64
	RandomString() string
	PublishRandomString(context.Context, uint64)
}

var _ GeneratorService = (*Generator)(nil)

type Generator struct {
	counter     *uint64
	EventSource EventService
}

func NewGenerator(start uint64) *Generator {
	return &Generator{
		counter: &start,
	}
}

func (g *Generator) Counter() uint64 {
	return atomic.AddUint64(g.counter, 1)
}

func (g *Generator) RandomString() string {
	return uuid.New().String()
}

func (g *Generator) PublishRandomString(ctx context.Context, uid uint64) {
	e := Event{Type: "random", Payload: "merhaba"}
	ticker := time.NewTicker(500 * time.Millisecond)

	for {
		select {
		case <-ctx.Done():
			log.Debug().Msg("context canceled")
			ticker.Stop()
			return
		case <-ticker.C:
			g.EventSource.PublishEvent(uid, e)

		}
	}

}
