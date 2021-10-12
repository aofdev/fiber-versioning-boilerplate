package adapters

import (
	"context"
	"fmt"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoAdapter interface {
	FindOne(filter bson.M) (interface{}, error)
	FindMany(filter bson.M) ([]interface{}, error)
	CloseMongoAdapter()
}

type mongoAdapter struct {
	connection *mongo.Client
	collection *mongo.Collection
}

func NewMongoAdapter(uri, database, collection string) MongoAdapter {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	clientOpts := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(ctx, clientOpts)

	if err != nil {
		panic(fmt.Errorf("fatal error mongo connection: %s", err.Error()))
	}

	err = client.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	return &mongoAdapter{client, client.Database(database).Collection(collection)}
}

func (m *mongoAdapter) FindOne(filter bson.M) (interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	var bsonDocument bson.M
	err := m.collection.FindOne(ctx, filter).Decode(&bsonDocument)

	return bsonDocument, err
}

func (m *mongoAdapter) FindMany(filter bson.M) ([]interface{}, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	var result []interface{}
	cursor, err := m.collection.Find(ctx, filter)
	for cursor.Next(context.TODO()) {
		var bsonDocument bson.M
		_ = cursor.Decode(&bsonDocument)
		result = append(result, bsonDocument)
	}
	return result, err
}

func (m *mongoAdapter) CloseMongoAdapter() {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	err := m.connection.Disconnect(ctx)
	if err != nil {
		panic(fmt.Errorf("fatal error disconnection mongo: %s", err.Error()))
	}
}
