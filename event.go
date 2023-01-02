package dezge

import "context"

type Listener interface {
	Listen()
	Subs() chan string
}

type Event struct {
	Type string `json:"type"`
	// The actual data from the event. See related payload types below.
	Payload interface{} `json:"payload"`
}

type Subscription interface {
	Close() error
	C() <-chan Event
}

type EventService interface {
	PublishEvent(int, Event)
	Subscribe(ctx context.Context) (Subscription, error)
	LenSubs() int
}
