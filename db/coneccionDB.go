package db

import (
	"context"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

/* MongoC is the conection object to DB */
var MongoCN = ConnectDB()

var clientOptions = options.Client().ApplyURI("mongodb+srv://dbusers:dbuser1983@cluster0.x9vms.mongodb.net/twitter-golang?retryWrites=true&w=majority")

/* ConnectDB allow DB conextion*/
func ConnectDB() *mongo.Client {
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err.Error())
		return client
	}
	log.Println("Conexion Exitosa")
	return client

}

/* ConectionCheck is ping ot DB*/
func ConectionCheck() int {
	err := MongoCN.Ping(context.TODO(), nil)
	if err != nil {
		return 0
	}
	return 1
}
