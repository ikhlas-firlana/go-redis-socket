package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/go-redis/redis"
	socketio "github.com/googollee/go-socket.io"
)

func handleRecover() {
	if err := recover(); err != nil {
		log.Print("NOT RIGHT!")
	}
}

func main() {
	defer handleRecover()
	host := os.Getenv("HOST_REDIS")
	portRedis := os.Getenv("PORT_REDIS")
	port := os.Getenv("PORT")
	client := redis.NewClient(&redis.Options{
		Addr:     host + ":" + portRedis,
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	log.Printf("host %v \n", host)
	log.Printf("portRedis %v \n", portRedis)
	log.Printf("port %v \n", port)
	log.Println("Working!")

	_, err := client.Ping().Result()
	if err != nil {
		panic(err)
	}
	server, err := socketio.NewServer(nil)
	if err != nil {
		panic(err)
	}
	//
	server.On("connection", func(so socketio.Socket) {
		log.Println("on connection")

		so.On("disconnection", func() {
			log.Println("on disconnect")
		})

		pubsub := client.Subscribe()

		go func() {
			defer handleRecover()
			err = pubsub.Subscribe("yuhu channel")
			if err != nil {
				panic(err)
			}
			for {
				msg, err := pubsub.ReceiveMessage()
				if err != nil {
					log.Println("err", fmt.Errorf("%v", err))
				}
				log.Println(msg.Channel, msg.Payload)
				// so.Emit("new_message", msg)
				log.Println("emit:", so.Emit("new_message", msg))
			}
		}()

	})
	server.On("error", func(so socketio.Socket, err error) {
		log.Println("error:", err)
		panic(err)
	})

	//
	http.Handle("/socket.io/", server)
	log.Println(http.ListenAndServe(":"+port, nil))
}
