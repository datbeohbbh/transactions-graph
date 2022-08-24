package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
)

func scaleMethod(ctx context.Context) (int, error) {
	if os.Getenv("HOSTNAME") == "" {
		return -1, errors.New("undefined host name")
	}

	hostname := os.Getenv("HOSTNAME")
	mongoConn, err := connectToMongoDB(os.Getenv("MONGODB_URI"), os.Getenv("MONGO_DATABASE"), os.Getenv("MONGO_USERNAME"), os.Getenv("MONGO_PASSWORD"))
	if err != nil {
		return -1, fmt.Errorf("failed on connect to mongo: %v", err)
	}
	defer mongoConn.Disconnect(ctx)

	type HostName struct {
		Name     string   `json:"name" bson:"name"`
		HostList []string `json:"hostList" bson:"hostList"`
	}

	scale := connectToDB(mongoConn, "graphdb")
	scaleColl := scale.Collection("scale")

	cnt, err := scaleColl.EstimatedDocumentCount(ctx)
	if err != nil {
		return -1, fmt.Errorf("failed on estimate the number of document")
	}

	if cnt == 0 {
		_, err := scaleColl.InsertOne(ctx, HostName{Name: "host", HostList: []string{}})
		if err != nil {
			return -1, fmt.Errorf("failed on insert initial document: %v", err)
		}
	}

	filter := bson.D{{"name", "host"}}
	push := bson.D{{"$push", bson.D{{"hostList", hostname}}}}

	_, err = scaleColl.UpdateOne(ctx, filter, push)
	if err != nil {
		return -1, fmt.Errorf("failed on push host name: %v", err)
	}

	hostList := HostName{}
	err = scaleColl.FindOne(ctx, filter).Decode(&hostList)
	if err != nil {
		return -1, fmt.Errorf("failed on get host list: %v", err)
	}

	log.Println("Getting Serving range")
	for i := len(hostList.HostList) - 1; i >= 0; i-- {
		if hostList.HostList[i] == hostname {
			return i % 5, nil
		}
	}
	return -1, nil
}
