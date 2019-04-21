package main

import (
	"fmt"
	"go-rabbitmq-test/distr_course/coordinator"
)

var dc *coordinator.DatabaseConsumer
var wc *coordinator.WebappConsumer

func main () {
	ea := coordinator.NewEventAggregator()
	dc = coordinator.NewDatabaseConsumer(ea)
	wc = coordinator.NewWebappConsumer(ea)
	ql:= coordinator.NewQueuelistener(ea)

	go ql.ListenForNewSource()

	var a string
	fmt.Scanln(&a)
}