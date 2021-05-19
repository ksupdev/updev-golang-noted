package main

import (
	"encoding/json"
	"fmt"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

/*
	Gorm auto generate table name form Product struct => products
	field "Code" generate to "code"
	field "Price" generate to "price"

	field gorm.Model (ID,CreatedAt,UpdatedAt,DeletedAt) => id,created_at,updated_at,deleted_at

	gorm.Model don't requir field

*/

type Product struct {
	gorm.Model // Use for genereate standare column ID,CreatedAt,UpdatedAt,DeletedAt
	Code       string
	Price      uint
}

func main() {
	db, err := gorm.Open(sqlite.Open("test.db"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	//Migrate the schema
	// Auto check schema if doesn't exit then it's auto generate schema
	db.AutoMigrate(&Product{})

	// Create
	db.Create(&Product{Code: "D42", Price: 100})

	// Read
	var product Product
	db.First(&product, 1) // Find product with integer primary key

	// log to console
	fmt.Println(product)
	json, _ := json.MarshalIndent(product, "", " ")
	fmt.Println(string(json))

	db.First(&product, "Code = ?", "D42") // find product with code D42

	// Update - update product's price to 200
	db.Model(&product).Update("Price", 200)
	// Update - update multiple fields
	db.Model(&product).Updates(Product{Price: 200, Code: "F42"}) // non-zero fields
	db.Model(&product).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	//Delete - delete product
	//db.Delete(&product, 1)

}
