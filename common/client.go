package common

import (
	"encoding/json"
	"fmt"
	"log"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

const Broker = "localhost:1891"
const IdPublisher = "go-mqtt-publisher"
const IdSubscriber = "go-mqtt-subscriber"

var Handler mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {

	CheckTemperature(msg.Payload())

	return
}

func CreateClient(broker string, id string, callback_handler mqtt.MessageHandler) mqtt.Client {

	opts := mqtt.NewClientOptions().AddBroker(broker)
	opts.SetClientID(id)
	opts.SetDefaultPublishHandler(callback_handler)

	return mqtt.NewClient(opts)
}

func CheckTemperature(payload []byte) {

	var data map[string]interface{}

	err := json.Unmarshal([]byte(string(payload)), &data)

	if err != nil {
		log.Fatalf("Algo deu errado na conversão dos dados...:", err)
	}

	id, ok := data["Id"].(string)
	if !ok {
		log.Fatalf("Failed to convert id to string")
	}

	s_type, ok := data["Sensor_type"].(string)
	if !ok {
		log.Fatalf("Failed to convert id to string")
	}

	temperature, ok := data["Temperature"].(float64)
	if !ok {
		log.Fatalf("Failed to convert temperature to float64")
	}

	fmt.Printf("%s: %s | %.0f°C", id, s_type, temperature)

	if s_type == "freezer" {
		if temperature > -15 {
			fmt.Println(" [ALERTA: Temperatura ALTA]")
		} else if temperature < -25 {
			fmt.Println(" [ALERTA: Temperatura BAIXA]")
		}
	} else {
		if temperature > 10 {
			fmt.Println(" [ALERTA: Temperatura ALTA]")
		} else if temperature < 2 {
			fmt.Println(" [ALERTA: Temperatura BAIXA]")
		}
	}

}
