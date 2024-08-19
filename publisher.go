package main

import (
	"encoding/json"
	"github.com/nats-io/nats.go"
	"io/ioutil"
	"log"
	"math/rand"
	"nats_jetstream/config"
	"nats_jetstream/models"
	"time"
)

func publishUsers(js nats.JetStreamContext) {
	users, err := getUsers()
	if err != nil {
		log.Println(err)
		return
	}

	for _, oneUser := range users {
		r := rand.Intn(2000)
		time.Sleep(time.Duration(r) * time.Millisecond)

		userString, err := json.Marshal(oneUser)

		if err != nil {
			log.Println(err)
			continue
		}

		_, err = js.Publish(config.SubjectNameUserCreated, userString)
		if err != nil {
			log.Println(err)
		} else {
			log.Printf("Publisher => Message: %s\n", oneUser.Name)
		}
	}
}

func getUsers() ([]models.User, error) {
	rawUsers, _ := ioutil.ReadFile("./users.json")
	var usersObj []models.User
	err := json.Unmarshal(rawUsers, &usersObj)

	return usersObj, err
}
