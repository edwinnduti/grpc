package natsConn

import (
	"os"

	"github.com/nats-io/nats.go"
)

func JsConnect() (nats.JetStreamContext, error) {
	// acquire nats connection
	nc, err := nats.Connect(os.Getenv("NATSURL"))
	if err != nil {
		return nil, err
	}

	// acquire jetstream context for messaging and stream management
	jetStream, err := nc.JetStream()
	if err != nil {
		return nil, err
	}
	return jetStream, nil
}