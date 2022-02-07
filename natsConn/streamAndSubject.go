package natsConn

import (
	"github.com/edwinnduti/natschat/logger"
	"github.com/nats-io/nats.go"
)

type JetStreamContext struct {
	Conn nats.JetStreamContext
}

func (js *JetStreamContext) NewStreamAndSubject(streamName string, subjectName string) error {
	// create a new stream
	stream, err := js.Conn.StreamInfo(streamName)
	if err != nil {
		return err
	}

	// if stream does not exist, create a new stream
	if stream == nil {
		logger.Success.Printf("Creating stream %s and subject %s", streamName, subjectName)
		_, err = js.Conn.AddStream(&nats.StreamConfig{
			Name:     streamName,
			Subjects: []string{subjectName},
		})
		return err
	}

	return nil
}