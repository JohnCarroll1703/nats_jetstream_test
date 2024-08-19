package main

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"log"
	"math/rand"
	"nats_jetstream/config"
	"nats_jetstream/models"
)

func consumeUsers(js nats.JetStreamContext) {
	_, err := js.Subscribe(config.SubjectNameUserCreated, func(msg *nats.Msg) {
		err := msg.Ack()

		if err != nil {
			log.Println("Error acknowledging message", err)
			return
		}

		var user models.User
		err = json.Unmarshal(msg.Data, &user)
		if err != nil {
			log.Println("Error unmarshalling user", err)
		}

		log.Printf("Consumer => Subject: %s - Id: %d - User: %s - Username: %s - Rand number: %o\n", msg.Subject,
			user.Id, user.Name, user.Username, rand.Intn(101))

		js.Publish(config.SubjectNameUserAnswered, []byte(user.Name))

	})

	if err != nil {
		log.Println("Subscribe failed")
		return
	}

}
