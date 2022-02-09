package database

import (
	"context"
	"fmt"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	ConnectionString string `mapstructure:"connection_string"`
	DBName           string `mapstructure:"database_name"`

	Client *mongo.Client
}

type Database struct {
	DB *mongo.Database
}

func (m *MongoDB) Binding() error {
	fmt.Println(m.ConnectionString)
	client, err := mongo.NewClient(options.Client().ApplyURI(m.ConnectionString))
	if err != nil {
		return err
	}
	if err := client.Connect(context.Background()); err != nil {
		return err
	}

	// check connection
	if err := client.Ping(context.Background(), nil); err != nil {
		return err
	}

	m.Client = client
	return nil
}

func NewDatabase(client *mongo.Client, dbName string) *Database {
	db := client.Database(dbName)
	return &Database{DB: db}
}
