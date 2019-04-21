package main

import (
	"go-rabbitmq-test/distr_course/web/controller"
	"net/http"
)

func main() {
	controller.Initialize()

	http.ListenAndServe(":3000", nil)
}