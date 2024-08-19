package main

import (
	"log"
	"sync"
)

func main() {
	log.Println("Starting JetStream...")

	js, err := JetStreamInit()

	if err != nil {
		log.Println(err)
		return
	}

	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		publishUsers(js)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		consumeUsers(js)
	}()

	wg.Wait()

	log.Println("Finished JetStream...")
}
