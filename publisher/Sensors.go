package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"time"

	DefaultClient "market/mqtt/common"
)

type Sensor struct {
	Id          string
	Sensor_type string
	Temperature int
	Time        string
}

func NewSensor(
	id string,
	sensor_type string,
	temperature int,
	time string) *Sensor {

	s := &Sensor{
		Id:          id,
		Sensor_type: sensor_type,
		Temperature: temperature,
		Time:        time,
	}

	return s

}

func (s *Sensor) ToJSON() (string, error) {
	jsonData, err := json.Marshal(s)
	if err != nil {
		return "", err
	}
	return string(jsonData), nil
}

func main() {
	freezer_sensor := NewSensor("lj01f01", "freezer", 0, "01/03/2024 14:30")

	client := DefaultClient.CreateClient(DefaultClient.Broker, DefaultClient.IdPublisher, DefaultClient.Handler)

	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	topic := "/" + freezer_sensor.Sensor_type

	for {

		payload, _ := freezer_sensor.ToJSON()

		freezer_sensor.Temperature = rand.Intn(40) - 30

		token := client.Publish(topic, 1, false, payload)

		token.Wait()

		fmt.Printf("Published message: %s\n", payload)

		time.Sleep(time.Duration(1) * time.Second)
	}

}
