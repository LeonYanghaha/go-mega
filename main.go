package main

import (
	"go-mega/controller"
	"go-mega/model"
	"net/http"
)

func main() {
	// Setup DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	// Setup Controller
	controller.Startup()
	http.ListenAndServe(":8888", nil)
}
