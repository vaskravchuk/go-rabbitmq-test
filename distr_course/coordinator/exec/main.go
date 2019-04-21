package main

import (
	"fmt"
	"go-rabbitmq-test/distr_course/coordinator"
)

func main () {
	ql:= coordinator.NewQueuelistener()

	go ql.ListenForNewSource()

	var a string
	fmt.Scanln(&a)
}