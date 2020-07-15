package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type Trainer struct {
	Title   string  `bson:"title"`
	Uuid    string  `bson:"uuid"`
	Content interface{} `bson:"content"`
}

func main() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017")

	// Connect to MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MongoDB!")

	// Get a handle for your collection
	collection := client.Database("test").Collection("mongotest")

	// Some dummy data to add to the Database
	//ash := Trainer{"Ash", "10", "Pallet Town"}
	//misty := Trainer{"Misty", "10", "Cerulean City"}
	//brock := Trainer{"Brock", "15", "Pewter City"}

	// Insert a single document
	//insertResult, err := collection.InsertOne(context.TODO(), ash)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Inserted a single document: ", insertResult.InsertedID)

	// Insert multiple documents
	//trainers := []interface{}{misty, brock}

	//insertManyResult, err := collection.InsertMany(context.TODO(), trainers)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println("Inserted multiple documents: ", insertManyResult.InsertedIDs)

	// Update a document
	filter := bson.D{{"uuid", "RC6fa883aa-c922-4a25-bcd1-428541fdede5"},{"title", "1212"}}


	//updateResult, err := collection.UpdateOne(context.TODO(), filter, update)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Printf("Matched %v documents and updated %v documents.\n", updateResult.MatchedCount, updateResult.ModifiedCount)

	// Find a single document
	var result Trainer

	err = collection.FindOne(context.TODO(), filter).Decode(&result)
	if err != nil {
		log.Fatal(err)
	}

	subMapB := make(map[string]interface{})
	var subMapB2  interface{}

	bb := result.Content.(bson.D)

	subMapB["name"] = result.Title
	subMapB["age"] = result.Uuid

	for _, v := range bb {
		for k1, v1 := range v.Value.(bson.D) {
			fmt.Println(k1, "   ", v1.Value)
		}

	}
	 dsd,err := collection.UpdateOne(context.TODO(), bson.M{"name": "howie_2"}, bson.M{"$set": bson.M{"name": "我要改了他的名字"}})
	fmt.Println(dsd)
	subMapB["content"] =subMapB2
	fmt.Printf("Found a single document: %+v\n", subMapB)

	//findOptions := options.Find()
	//findOptions.SetLimit(29)

	//var results []*Trainer
	//
	//// Finding multiple documents returns a cursor
	//cur, err := collection.Find(context.TODO(), bson.D{{}})
	//if err != nil {
	//	log.Fatal(err)
	//}
	//countryCapitalMap := make(map[int]map[string]string)
	//subMapB := make(map[string]string)
	//var ii int = 0;
	//
	//// Iterate through the cursor
	//
	//for cur.Next(context.TODO()) {
	//	var elem Trainer
	//	err := cur.Decode(&elem)
	//	if err != nil {
	//		log.Fatal(err)
	//	}
	//	results = append(results, &elem)
	//	//t := reflect.TypeOf(elem)
	//
	//}
	//for _, value := range results {
	//	subMapB["name"] = value.Name
	//	subMapB["age"] = value.Age
	//	subMapB["city"] = value.City
	//	ii++
	//	countryCapitalMap[ii] = subMapB
	//}
	//
	//fmt.Println("=====",countryCapitalMap)
	//if err := cur.Err(); err != nil {
	//	log.Fatal(err)
	//}

	// Close the cursor once finished
	//cur.Close(context.TODO())

	// Delete all the documents in the collection
	//deleteResult, err := collection.DeleteMany(context.TODO(), bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}

	//fmt.Printf("Deleted %v documents in the trainers collection\n", deleteResult.DeletedCount)

	// Close the connection once no longer needed
	err = client.Disconnect(context.TODO())

	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Connection to MongoDB closed.")
	}

}
