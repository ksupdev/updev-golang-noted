package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	var db *gorm.DB = connectDatabase()
	todoList := todov1{Username: "karoon", Title: "New developer", Message: "Hi how are you"}

	// Insert

	fmt.Println("---- Insert action ----")
	db.Create(&todoList)

	// Query
	fmt.Println("----- Query ---- ")
	query(db)

	// Update
	fmt.Println("---- Update ----")
	var tempForUpdate todov1
	db.First(&tempForUpdate)
	update(db, tempForUpdate)
	query(db)

	// Delete
	fmt.Println("---- Delete ----")
	var tempForDelete todov1
	db.Find(&tempForDelete, "message LIKE ?", "%%")
	delete(db, tempForDelete)
	query(db)
}

func delete(_db *gorm.DB, todo todov1) {
	_db.Unscoped().Delete(&todo)
}

func query(_db *gorm.DB) {
	var todos []todov1
	_db.Find(&todos)
	printPretty(todos)

}

func update(_db *gorm.DB, todos todov1) {
	_db.Model(&todos).Update("Message", "Last update")
}

func printPretty(data []todov1) {
	json, _ := json.MarshalIndent(data, "", " ")
	fmt.Println(string(json))
}

func connectDatabase() *gorm.DB {
	dsn := "root:rootpw@tcp()/cabin?parseTime=true&loc=Local"
	database, _ := gorm.Open(mysql.Open(dsn), &gorm.Config{})
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
	Username string `json:"username"`
	Title    string `json:"title"`
	Message  string `json:"message"`
}
