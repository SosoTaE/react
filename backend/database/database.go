package database

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	Client *mongo.Client
	Database mongo.Database 
	Collection mongo.Collection
}



func InitializeDatabase(connectionString, dbName, collName string) Database {
	client, err := mongo.Connect(context.TODO(), options.Client().ApplyURI(connectionString))
	if err != nil {
		panic(err)
	}

	db := client.Database(dbName)
	coll := db.Collection(collName)

	return Database{
		Client: client,
		Database: *db,
		Collection: *coll,
	}
}


