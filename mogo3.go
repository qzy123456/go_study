package main

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"time"
)

type Trainer1 struct {
	Name string
	Age  string
	City string
}

func main() {

	// Set client options
	clientOptions := options.Client().ApplyURI("mongodb://root:root@localhost:27017")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
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
	collection := client.Database("test").Collection("trainers")


	//findOptions := options.Find()
	//findOptions.SetLimit(29)

	//var results []*Trainer1

	// Finding multiple documents returns a cursor
	cur, err := collection.Find(ctx, bson.D{{}})
	if err != nil {
		log.Fatal(err)
	}
	countryCapitalMap := []map[string]interface{}{}
	subMapB := map[string]interface{}{}
	var ii int = 0;

	// Iterate through the cursor
	defer cur.Close(ctx)
	//var elem bson.M
	s := &Trainer1{}
	for cur.Next(ctx) {


		var result bson.M
		err := cur.Decode(&result) // decode 到map
		err = cur.Decode(s)  // decode 到对象
		if err != nil {
			log.Fatal(err)
		}

		// do something with result....
		// 可以将map 或对象序列化为json
		//js ,_:=json.Marshal(result)
		//json.Unmarshal(js,s) //反学序列化回来
		//for _, value := range result {
			//fmt.Println(s)
			subMapB["name"] = s.Name
			subMapB["age"] = s.Age
			subMapB["city"] = s.City
			ii++
			countryCapitalMap = append(countryCapitalMap,subMapB)

		//}
	}

fmt.Println(s)

	fmt.Println("=====",countryCapitalMap)
	if err := cur.Err(); err != nil {
		log.Fatal(err)
	}




}
