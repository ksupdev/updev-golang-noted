package store

import (
	"go.mongodb.org/mongo-driver/mongo"
	"updev.labs/up-order-service/order"
)

type MongoDBStoreMock struct {
	*mongo.Collection
}

func NewMongoDBStoreMock(dsn string) *MongoDBStoreMock {
	// client, err := mongo.Connect(context.Background(), options.Client().ApplyURI(dsn))
	// if err != nil {
	// 	panic("failed to connect database")
	// }
	// collection := client.Database("myapp").Collection("orders")
	return &MongoDBStoreMock{Collection: nil}
}

func (s *MongoDBStoreMock) Save(order order.Order) error {
	// _, err := s.Collection.InsertOne(context.Background(), order)
	return nil
}
