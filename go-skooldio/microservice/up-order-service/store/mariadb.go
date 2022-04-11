package store

import (
	"gorm.io/gorm"
	"updev.labs/up-order-service/order"
)

type GormStoreMock struct {
	db *gorm.DB
}

func NewMariaDBStoreMock(dsn string) *GormStoreMock {
	// db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	// if err != nil {
	// 	panic("failed to connect database")
	// }

	// db.AutoMigrate(&order.Order{})

	return &GormStoreMock{db: nil}
}

func (s *GormStoreMock) Save(o order.Order) error {
	return nil
}
