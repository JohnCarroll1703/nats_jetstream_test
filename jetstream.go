package main

import (
	"github.com/nats-io/nats.go"
	"log"
	"nats_jetstream/config"
)

func CreateStream(js nats.JetStreamContext) error {
	stream, err := js.StreamInfo(config.StreamName)
	if stream == nil {
		log.Printf("Creating stream: %s\n", config.StreamName)
		_, err = js.AddStream(&nats.StreamConfig{
			Name:     config.StreamName,
			Subjects: []string{config.StreamSubjects},
		})
		if err != nil {
			return err
		}
	}
	return nil
}

func JetStreamInit() (nats.JetStreamContext, error) {
	ns, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		return nil, err
	}

	js, err := ns.JetStream(nats.PublishAsyncMaxPending(256))
	if err != nil {
		return nil, err
	}

	err = CreateStream(js)
	if err != nil {
		return nil, err
	}

	return js, nil
}
