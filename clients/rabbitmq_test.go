package clients

import (
	"testing"
	"time"
)

// TestRabbitMQClient_Publish tests publishing messages to RabbitMQ
func TestRabbitMQClient_Publish(t *testing.T) {
	// このテストはdocker-composeでRabbitMQが起動している必要があります
	// docker-compose up -d rabbitmq

	// RabbitMQの起動を待つ
	time.Sleep(2 * time.Second)

	client, err := NewRabbitMQClient("amqp://guest:guest@localhost:17005/")
	if err != nil {
		t.Skipf("RabbitMQ not available: %v", err)
		return
	}
	defer client.Close()

	message := []byte(`{"type":"test","message":"Hello from test"}`)
	err = client.Publish("test_queue", message)
	if err != nil {
		t.Fatalf("Failed to publish message: %v", err)
	}

	t.Log("Successfully published message to test_queue")
}
