package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	client()
//	server()
}

func server() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	var input string
	for input != "q"  {
		fmt.Scanln(&input)

		msg := amqp.Publishing{
			ContentType: "text/plain",
			Body: []byte(input),
		}
		ch.Publish("", q.Name, false, false, msg)
	}
}

func client() {
	conn, ch, q := getQueue()
	defer conn.Close()
	defer ch.Close()

	msgs, err := ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)
	failOnError(err, "Failed to create consumer")

	for msg := range msgs {
		log.Printf("Received message with message: %s", msg.Body)
	}
}

func getQueue() (*amqp.Connection, *amqp.Channel, *amqp.Queue) {
	conn, err := amqp.Dial("amqp://guest@10.0.1.13:5672")
	failOnError(err, "Failed to connect to RabbitMQ")
	ch, err := conn.Channel()
	failOnError(err, "Failed to open a Channel")
	q, err := ch.QueueDeclare("hello", false, false, false, false, nil)
	failOnError(err, "Failed to open declare a queue")

	return conn, ch, &q
}

func failOnError(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %s", msg, err)
		panic(fmt.Sprintf("%s: %s", msg, err))
	}
}
