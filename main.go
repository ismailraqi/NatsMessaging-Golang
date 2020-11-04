package main

import (
	"log"
	"time"

	"github.com/nats-io/nats.go"
)

func main() {
	// Connecting to the Default Server
	nc, err := nats.Connect(nats.DefaultURL,
		// Set max reconnects attempts
		nats.Name("natsMessaging"), nats.MaxReconnects(10),
		// Ping/Pong Protocol
		nats.MaxPingsOutstanding(5), nats.PingInterval(20*time.Second))
	//nats.UserInfo("ismail", "password123"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()
	// Maximum Payload Size
	mp := nc.MaxPayload()
	log.Printf("Maximum payload is %v bytes", mp)
	// sending message
	if err := nc.Publish("updates", []byte("All is Well")); err != nil {
		log.Fatal(err)
	}
	// sending message with reply
	if err := nc.Publish("updates", []byte("All is Well")); err != nil {
		log.Fatal(err)
	}
	// Sync Subscription
	sub, err := nc.SubscribeSync("updates")
	if err != nil {
		log.Fatal(err)
	}
	// Read a message
	msg, err := sub.NextMsg(10 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// Get the time
	timeAsBytes := []byte(time.Now().String())

	// Send the time as the response.
	msg.Respond(timeAsBytes)
	if err := sub.Unsubscribe(); err != nil {
		log.Fatal(err)
	}

	// Async Subscription
	sub, err = nc.Subscribe("updates", func(_ *nats.Msg) {})
	if err != nil {
		log.Fatal(err)
	}
	// Read a message
	msg2, err := sub.NextMsg(10 * time.Second)
	if err != nil {
		log.Fatal(err)
	}

	// Get the time
	time2AsBytes := []byte(time.Now().String())

	// Send the time as the response.
	msg2.Respond(time2AsBytes)
	// Unsubscribing After N Messages
	if err := sub.AutoUnsubscribe(1); err != nil {
		log.Fatal(err)
	}
}
