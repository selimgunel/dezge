package rabbitmq

import (
	"fmt"

	"github.com/narslan/dezge"
	"github.com/streadway/amqp"
)

var _ dezge.Listener = (*Listener)(nil)

type Listener struct {
	writeChan chan string
}

func NewListener() *Listener {
	return &Listener{
		writeChan: make(chan string),
	}
}

func (l *Listener) Listen() {
	conn, err := amqp.Dial("amqp://user:bitnami@localhost:5672/")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer conn.Close()
	ch, err := conn.Channel()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer ch.Close()
	msgs, err := ch.Consume(
		"test_que",               // queue
		"SELAM BEN BİR CONSUMER", // consumer tag
		true,                     // auto-ack
		false,                    // exclusive
		false,                    // no-local
		false,                    // no-wait
		nil,                      // args
	)
	if err != nil {
		fmt.Println(err)
		return
	}

	forever := make(chan bool)
	go func() {
		for d := range msgs {
			l.writeChan <- string(d.Body)
		}
	}()
	<-forever

}

func (l *Listener) Subs() chan string {
	return l.writeChan
}
