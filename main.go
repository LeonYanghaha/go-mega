package main

import (
	"github.com/gorilla/context"
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
	_ = http.ListenAndServe(":8888", context.ClearHandler(http.DefaultServeMux))
}
