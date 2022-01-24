package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/go-redis/redis/v8"
)

type Details struct {
	Message string `json:"message"`
}

func (details *Details) toString() string {
	return fmt.Sprintf("message=%s", details.Message)
}

type Response struct {
	Message string `json:"message"`
}

func loginHandler() {
	redisMessage := &Details{}

	var redisClient = redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})

	subscriber := redisClient.Subscribe(context.Background(), "send-idle-notification-message")

	for {
		msg, err := subscriber.ReceiveMessage(context.Background())
		if err != nil {
			panic(err)
		}

		if err := json.Unmarshal([]byte(msg.Payload), &redisMessage); err != nil {
			panic(err)
		}

		fmt.Println("Message received from " + msg.Channel + " channel: => " + redisMessage.toString())
	}
}

func main() {
	loginHandler()
}
