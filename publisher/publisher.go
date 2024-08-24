package publisher

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/brianvoe/gofakeit/v6"
	"github.com/nats-io/nats.go"
)

type FakeData struct {
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}

func Publish() {
	// Connect to the NATS server
	nc, err := nats.Connect(nats.DefaultURL)
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	// Subject to which we will publish the data
	subject := "test.subject"
	// Generate and publish fake data
	for i := 0; i < 10; i++ {
		// Generate fake data

		name := gofakeit.Name()
		email := gofakeit.Email()
		phone := gofakeit.Phone()

		data := FakeData{
			Name:  name,
			Email: email,
			Phone: phone,
		}

		// Serialize the data to JSON
		jsonData, err := json.Marshal(data)
		if err != nil {
			log.Fatal(err)
		}

		// Publish the data to the NATS subject
		err = nc.Publish(subject, jsonData)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Published message %d: %s\n", i+1, jsonData)

		// Sleep for a short time before sending the next message
		time.Sleep(1 * time.Second)
	}

	// Ensure the message has been sent
	nc.Flush()

	// Check for any errors during publishing
	if err := nc.LastError(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("All messages published successfully!")
}
