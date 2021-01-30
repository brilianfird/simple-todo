package config

import (
	"context"
	"github.com/tkanos/gonfig"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

var DB *mongo.Client

var Database *mongo.Database

func InitDb() {
	prop := MongoProperties{}
	gonfig.GetConf("conf.json", &prop)
	log.Println("MONGO URI: ", prop.URI)
	client, err := mongo.NewClient(options.Client().ApplyURI(prop.URI))

	if err != nil {
		log.Fatal("Error initializing mongodb", err)
	}
	ctx, _ := context.WithTimeout(context.Background(), 10 * time.Second)

	client.Connect(ctx)

	DB = client
	Database = DB.Database(prop.DatabaseName)
}