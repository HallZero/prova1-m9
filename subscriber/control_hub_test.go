package main

import (
	"testing"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	DefaultClient "market/mqtt/common"
)

var subscriber = DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdSubscriber, DefaultClient.Handler)

func TestPublisher(t *testing.T) {
	t.Run("Test QoS - eg if the message was received by the broker and testing alarms", func(t *testing.T) {

		if token := subscriber.Connect(); token.Wait() && token.Error() != nil {
			t.Fatal(token.Error())
		}

		defer subscriber.Disconnect(250)

		topic := "#"

		received := make(chan []byte)

		subscriber.Subscribe(topic, 1, func(client mqtt.Client, message mqtt.Message) {
			
			DefaultClient.CheckTemperature(message.Payload())
			
			received <- message.Payload()
		})

		// Publish a message
		message := "test payload"
		subscriber.Publish(topic, 1, false, message)

		// Wait for a short duration to receive the message
		select {
		case payload := <-received:
			if string(payload) != message {
				t.Errorf("Received payload %s, expected %s", payload, message)
			}
		case <-time.After(2 * time.Second):
			t.Error("Timeout: Did not receive the payload")
		}

	})

	// t.Run("Testing alarms", func(t *testing.T) {
	// 	if token := subscriber.Connect(); token.Wait() && token.Error() != nil {
	// 		t.Fatal(token.Error())
	// 	}

	// 	defer subscriber.Disconnect(250)

	// 	topic := "#"


	// })
}
