package main

import (
	"encoding/json"
	"testing"
	"fmt"

	DefaultClient "market/mqtt/common"
)

var client = DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdPublisher, DefaultClient.Handler)

func TestPublisher(t *testing.T) {

	t.Run("Generating JSON file to payload", func(t *testing.T) {
		sensor := NewSensor("lj01f01", "freezer", -18, "01/03/2024 14:30")

		got, err := sensor.ToJSON()

		var transformed map[string]interface{}

		json.Unmarshal([]byte(got), &transformed)

		if err != nil {
			t.Errorf("Error generating JSON: %v", err)
		}

		want := map[string]interface{}{
			"Id":          "lj01f01",
			"Sensor_type":    "freezer",
			"Temperature":   -18,
			"Time": "01/03/2024 14:30",
		}

		// May change this later. Map comparison is quite confusing (reflect.DeepEqual() returns false)
		if !(fmt.Sprint(transformed) == fmt.Sprint(want)) {
			t.Errorf("Unexpected JSON output.\nGot: %v\nWant: %v", transformed, want)
		}

	})

	t.Run("Test QoS - eg if the message was published by the broker", func(t *testing.T) {

		payload := "Hello, Broker!"

		if token := client.Connect(); token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}

		token := client.Publish("sensors", 1, false, payload)

		if token.Wait() && token.Error() != nil {
			t.Error(token.Error())
		}

		t.Log("Broker received message with QoS 1!")
	})
}
