package model

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"log"
	"mega/config"
)

var db *gorm.DB

// SetDB func
func SetDB(database *gorm.DB) {
	db = database
}

// ConnectToDB func
func ConnectToDB() *gorm.DB {
	connectingStr := config.GetMysqlConnectingString()
	log.Println("Connet to db...")
	fmt.Println(connectingStr)
	//dbs:="root:123@tcp(localhost:3306)/mega?charset=utf8&parseTime=True&loc=Local"
	db, err := gorm.Open("mysql", connectingStr)
	if err != nil {
		panic("Failed to connect database")
	}
	db.SingularTable(true)
	return db
}