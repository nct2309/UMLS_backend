package driver

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

type MongoDB struct {
	Client *mongo.Client
}

var Mongo = &MongoDB{}

func ConnectMongoDB() *MongoDB {

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb+srv://kienle123:123456789vaythoi@cluster0.3fp6qr3.mongodb.net/"))
	if err != nil {
		panic(err)

	}
	err = client.Ping(ctx, readpref.Primary()) // ping bản chính, not secondary
	if err != nil {
		panic(err)
	}

	fmt.Println("connection is ok now")
	Mongo.Client = client
	return Mongo
}
