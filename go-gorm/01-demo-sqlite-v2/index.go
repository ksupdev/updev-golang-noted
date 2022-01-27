package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB = connectDatabase()
	todoList := todov1{Username: "admin", Title: "task1", Message: "Message 1"}

	//Create
	fmt.Println("---- create ----")
	db.Create(&todoList)

	//Update
	fmt.Println("---- update ----")
	var tmpForUpdate []todov1
	db.First(&tmpForUpdate, 1)
	update(db, tmpForUpdate)

	//Query
	fmt.Println("---- query ----")
	query(db)

	//Delete
	fmt.Println("---- delete ----")
	var tmpForDelete []todov1
	db.Find(&tmpForDelete, "message LIKE ?", "%updev%")
	delete(db, tmpForDelete)
	query(db)
}

func delete(_db *gorm.DB, todos []todov1) {
	// Hard delete delete value
	if len(todos) != 0 {
		_db.Unscoped().Delete(&todos)
	}
	// Soft delete mark field DeleteAt with Delete date
	// _db.Delete(&todos)
}

func query(_db *gorm.DB) {
	var todos []todov1
	_db.Find(&todos)
	printPretty(todos)

}

func update(_db *gorm.DB, todos []todov1) {
	if len(todos) != 0 {
		_db.Model(&todos).Update("Message", "update updev")
	}

}

func printPretty(data interface{}) {
	json1, _ := json.MarshalIndent(data, "", " ")
	fmt.Println(string(json1))
}

func connectDatabase() *gorm.DB {
	database, _ := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	database.AutoMigrate(&todov1{})
	return database

}

type todov1 struct {
	gorm.Model
	Username string
	Title    string
	Message  string
}
type todov2 struct {
	ID       uint   `json:"id" gorm:"primary_key"`
	Username string `json:"username`
	Title    string `json:"title"`
	Message  string `json:"message"`
}
