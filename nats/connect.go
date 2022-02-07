package natsConn

import (
	"github.com/nats-io/nats.go"
)

func JsConnect() (nats.JetStreamContext, error) {
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}
	jetStream, err := nc.JetStream()
	if err != nil {
		return nil, err
	}
	return jetStream, nil
}