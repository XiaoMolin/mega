package main

import (
	"github.com/gorilla/context"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"mega/controller"
	"mega/model"
	"net/http"
)

func main() {
	// Setup DB
	db := model.ConnectToDB()
	defer db.Close()
	model.SetDB(db)

	// Setup Controller
	controller.Startup()

	http.ListenAndServe(":8888", context.ClearHandler(http.DefaultServeMux))
}