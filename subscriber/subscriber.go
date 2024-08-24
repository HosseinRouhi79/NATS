package subscriber

import (
	"fmt"
	"log"

	"github.com/nats-io/nats.go"
)

func Subscribe() {
	// Connect to the NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subject to subscribe to
	subject := "test.subject"

	fmt.Println("test")
	// Subscribe to the subject
	_, err = nc.Subscribe(subject, func(msg *nats.Msg) {
		// Handle the message
		fmt.Printf("Received message: %s\n", string(msg.Data))
	})
	if err != nil {
		fmt.Printf("error:%s", err.Error())
	}

	// Wait for messages
	select {}
}
