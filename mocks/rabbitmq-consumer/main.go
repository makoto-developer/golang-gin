package main

import (
	"log"
	"time"

	amqp "github.com/rabbitmq/amqp091-go"
)

func main() {
	// RabbitMQへの接続を待つ
	time.Sleep(10 * time.Second)

	conn, err := amqp.Dial("amqp://guest:guest@rabbitmq:5672/")
	if err != nil {
		log.Fatalf("Failed to connect to RabbitMQ: %v", err)
	}
	defer conn.Close()

	ch, err := conn.Channel()
	if err != nil {
		log.Fatalf("Failed to open a channel: %v", err)
	}
	defer ch.Close()

	// キュー宣言
	q, err := ch.QueueDeclare(
		"test_queue", // name
		false,        // durable
		false,        // delete when unused
		false,        // exclusive
		false,        // no-wait
		nil,          // arguments
	)
	if err != nil {
		log.Fatalf("Failed to declare a queue: %v", err)
	}

	msgs, err := ch.Consume(
		q.Name, // queue
		"",     // consumer
		true,   // auto-ack
		false,  // exclusive
		false,  // no-local
		false,  // no-wait
		nil,    // args
	)
	if err != nil {
		log.Fatalf("Failed to register a consumer: %v", err)
	}

	log.Println("🚀 RabbitMQ Consumer Mock starting...")
	log.Println("📨 Waiting for messages on queue: test_queue")

	forever := make(chan bool)

	go func() {
		for d := range msgs {
			log.Printf("📬 Received message: %s", d.Body)
			log.Printf("   Content-Type: %s", d.ContentType)
			log.Printf("   Delivery Tag: %d", d.DeliveryTag)
		}
	}()

	<-forever
}
