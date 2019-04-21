package main

import (
	"fmt"
	"go-rabbitmq-test/distr_course/coordinator"
)

var dc *coordinator.DatabaseConsumer

func main () {
	ea := coordinator.NewEventAggregator()
	dc = coordinator.NewDatabaseConsumer(ea)
	ql:= coordinator.NewQueuelistener(ea)

	go ql.ListenForNewSource()

	var a string
	fmt.Scanln(&a)
}