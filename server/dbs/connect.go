package dbs

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	Uri             string = "mongodb://localhost:27017/"
	DbName          string = "VueLinkTree"
	UserCollection  *mongo.Collection
	AdminCollection *mongo.Collection
)

func init() {
	var clientOptions = options.Client().ApplyURI(Uri)

	var connected, conectionErr = mongo.Connect(context.TODO(), clientOptions)
	if conectionErr != nil {
		log.Panicln("Db connection Err")
		return
	}
	UserCollection = connected.Database(DbName).Collection("Users")
	AdminCollection = connected.Database(DbName).Collection("Admin")
}
