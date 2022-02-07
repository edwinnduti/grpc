package main

import (
	"log"
	"time"

	"github.com/edwinnduti/natschat/natsConn"
	"github.com/nats-io/nats.go"
)

var (
	streamName  = "ORDERS"
	subjectName = "ORDERS.received"
)

func main(){
	js, _ := natsConn.JsConnect()
	stream, err := js.StreamInfo("foo")
	if err != nil {
		log.Print(err)
	}

	if stream == nil {
		log.Printf("creating stream %q and subject %q", streamName, subjectName)
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{subjectName},
		})
		if err != nil {
			log.Print(err)
		}
	}

	log.Print("publishing an order")
	_, err = js.Publish(subjectName, []byte("one big burger"))
	if err != nil {
		log.Print(err)
	}

	log.Print("attempting to receive order")
	var order []byte
	done := make(chan bool, 1)
	js.Subscribe(subjectName, func(m *nats.Msg) {
		order = m.Data
		m.Ack()
		done <- true
	})

	select {
	case <-time.After(5 * time.Second):
		log.Fatalf("failed to get order")
	case <-done:
		log.Printf("got order: %q", order)
	}
}