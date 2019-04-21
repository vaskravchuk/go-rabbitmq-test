package coordinator

import (
	"bytes"
	"encoding/gob"
	"fmt"
	"github.com/streadway/amqp"
	"go-rabbitmq-test/distr_course/dto"
	"go-rabbitmq-test/distr_course/qutils"
)

const url = "amqp://guest:guest@10.0.1.13:5672"


type QueueListener struct {
	conn *amqp.Connection
	ch   *amqp.Channel
	sources map[string]<- chan amqp.Delivery
	ea *EventAggregator
}

func NewQueuelistener() *QueueListener {
	ql := QueueListener {
		sources:make(map[string]<- chan amqp.Delivery),
		ea: NewEventAggregator(),
	}

	ql.conn, ql.ch = qutils.GetChannel(url)

	return &ql
}

func (ql *QueueListener)DiscoverSensors() {
	ql.ch.ExchangeDeclare(
		qutils.SensorDiscoveryExchange,
		"fanout",
		false,
		false,
		false,
		false,
		nil)

	ql.ch.Publish(
		qutils.SensorDiscoveryExchange,
		"",
		false,
		false,
		amqp.Publishing{})
}

func (ql *QueueListener) ListenForNewSource() {
	q := qutils.GetQueue("", ql.ch)
	ql.ch.QueueBind(q.Name, "", "amq.fanout", false, nil)

	msgs, _ :=  ql.ch.Consume(
		q.Name,
		"",
		true,
		false,
		false,
		false,
		nil)

	ql.DiscoverSensors()

	for msg := range msgs {
		sourceChan, _ := ql.ch.Consume(
			string(msg.Body),
			"",
			true,
			false,
			false,
			false,
			nil)

		if ql.sources[string(msg.Body)] == nil {
			ql.sources[string(msg.Body)] = sourceChan

			go ql.AddListener(sourceChan)
		}
	}
}

func (ql *QueueListener) AddListener(msgs <- chan amqp.Delivery) {
	for msg := range msgs {
		r := bytes.NewReader(msg.Body)
		d := gob.NewDecoder(r)
		sd := new(dto.SensorMessage)
		d.Decode(sd)

		fmt.Printf("Received message %v\n", sd)

		ed := EventData{
			Name:      sd.Name,
			Value:     sd.Value,
			Timestamp: sd.Timestamp,
		}

		ql.ea.PublishEvent("MessageReceived_" + msg.RoutingKey, ed)
	}
}