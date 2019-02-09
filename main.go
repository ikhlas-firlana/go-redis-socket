package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-redis/redis"
	"github.com/googollee/go-socket.io"
)

func main() {
	client := redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}

	pubsub := client.Subscribe()

	go func() {
		err = pubsub.Subscribe("yuhu channel")
		if err != nil {
			panic(err)
		}

		log.Println("Stand by listener!")
		for {
			msg, err := pubsub.ReceiveMessage()
			if err != nil {
				log.Println("err", fmt.Errorf("%v", err))
			}
			log.Println(msg.Channel, msg.Payload)
		}
	}()

	log.Fatal(http.ListenAndServe(":5000", nil))
}
