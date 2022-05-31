package database

import (
	"context"
	"log"
	"time"

	"github.com/ayo-ajayi/gqlgen-todos/graph/model"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"github.com/ayo-ajayi/selfGin/config"
)
type DB struct{
	client *mongo.Client
}
var _, mongo_url, _ = config.Config("", "MONGODB_URL", "")
func Connect()*DB{
	client, err := mongo.NewClient(options.Client().ApplyURI(mongo_url))
	if err != nil {
		log.Fatal(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), 100*time.Second)
	defer cancel()
	err = client.Connect(ctx)
	if err != nil {
		log.Fatal(err)
	}
	return &DB{
		client: client,
	}
}
func(db* DB)Save(input *model.NewDog)*model.Dog{
	collection:=db.client.Database("animals").Collection("dogs")
	ctx, cancel:=context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res, err := collection.InsertOne(ctx, input)
	if err != nil {
		log.Fatal(err)
	}
	return &model.Dog{
		ID:        res.InsertedID.(primitive.ObjectID).Hex(),
		Name:      input.Name,
		IsGoodBoi: input.IsGoodBoi,
	}
}

func(db* DB)FindByID(ID string)*model.Dog{
	ObjectID, err := primitive.ObjectIDFromHex(ID)
	if err != nil {
		log.Fatal(err)
	}
	collection:=db.client.Database("animals").Collection("dogs")
	ctx, cancel:=context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	res:= collection.FindOne(ctx, bson.M{"_id": ObjectID})
	if err != nil {
		log.Fatal(err)
	}
	dog:=model.Dog{}
	res.Decode(&dog)
	return &dog
}

func(db* DB)All()[]*model.Dog{
	collection:=db.client.Database("animals").Collection("dogs")
	ctx, cancel:=context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	cur, err:=collection.Find(ctx, bson.D{})
	if err != nil {
		log.Fatal(err)
	}
	var dogs []*model.Dog
	for cur.Next(ctx){
		var dog *model.Dog
		err := cur.Decode(&dog)
		if err != nil {
			log.Fatal(err)
		}
		dogs = append(dogs, dog)
	}
	return dogs
}

