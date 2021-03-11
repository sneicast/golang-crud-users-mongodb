package database

import (
	"context"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	usr      = "usuarioDB"
	pwd      = "ContraseñaDB"
	host     = "localhost:27017"
	database = "nombreDB"
)

//GetCollection Collection
func GetCollection(collection string) *mongo.Collection {

	//uri := fmt.Sprintf("mongodb://%s:%s@%s:%d", usr, pwd, host, port)
	uri := fmt.Sprintf("mongodb+srv://%s:%s@%s", usr, pwd, host)

	client, err := mongo.NewClient(options.Client().ApplyURI(uri))

	if err != nil {
		panic(err.Error())
	}

	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	err = client.Connect(ctx)

	if err != nil {
		panic(err.Error())
	}

	return client.Database(database).Collection(collection)
}
