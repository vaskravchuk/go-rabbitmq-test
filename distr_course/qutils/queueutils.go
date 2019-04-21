package qutils

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

const (
	SensorDiscoveryExchange = "SensorDiscovery"
	PersistReadingsQueue = "PersistReadings"
)

func GetChannel(url string) (*amqp.Connection, *amqp.Channel) {
	conn, err := amqp.Dial(url)
	failOnError(err, "Failed to establisj connection to message broker")
	ch, err := conn.Channel()
	failOnError(err, "Failed to get Channel for connection")

	return conn, ch
}

func GetQueue(name string, ch *amqp.Channel, autoDelete bool) *amqp.Queue {
	q, err := ch.QueueDeclare(
		name,
		false,
		autoDelete,
		false,
		false,
		nil)
	failOnError(err, "Failed to declare queue")
	return &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}