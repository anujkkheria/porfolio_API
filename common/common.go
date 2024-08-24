package common

import (
	"context"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	// "go.mongodb.org/mongo-driver/options"
)

var db *mongo.Database

func GetDBCollection(col string)*mongo.Collection{
	return	db.Collection(col)
}

func InitDB() error{
	// uri:=os.Getenv("MONGODB_URI")
	// if uri == " "{
	// 	return errors.New("db uri missing")
	// }
	client, err := mongo.Connect(context.Background(), options.Client().ApplyURI("mongodb+srv://ankurguma:Wmq46HCJpNZGXLbZ@cluster0.jk9vz.mongodb.net/?retryWrites=true&w=majority&appName=Cluster0"))

	if err!= nil{
		return err
	}

	db = client.Database("portfolio")

	return nil
}

func CloseDB() error{

	return db.Client().Disconnect(context.Background())

}