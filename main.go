package main

import (
	"go-mega/controller"
	"net/http"
)

func main() {
	controller.Startup()
	http.ListenAndServe(":8888", nil)
}
