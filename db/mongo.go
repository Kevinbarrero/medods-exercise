package db

import (
	"context"
	"log"
	"medods-exercise/utils"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func Setup() *mongo.Database {

	config, err := utils.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config:", err)
	}
	clientOps := options.Client().ApplyURI(config.DBSOURCE)
	conn, err := mongo.Connect(context.Background(), clientOps)
	db := conn.Database("medods")
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	err = db.CreateCollection(context.Background(), "users")
	if err != nil {
		log.Fatal("cannot create collection", err)
	}
	return db
}
