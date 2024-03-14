package main

import (
	"context"
	node "einsof/sauron/pkg"
	"fmt"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	config *Config
	client *mongo.Client
}

func (db *Database) Init(config *Config) error {
  var err error
	
  ctx, cancel := context.WithTimeout(
		context.Background(), 20*time.Second,
	)
  defer cancel()
  
	db.config       = config
	db.client, err  = mongo.Connect(
		ctx,
		options.Client().ApplyURI(
			fmt.Sprintf(
				"mongodb://%s:%s@%s:%s",
				config.DBUsername,
				config.DBPassword,
				config.DBHost,
				config.DBPort,
			),
		),
	)

	if err != nil {
		fmt.Println(err)
	}

	return err
}

// TODO: make collections enum & refactor to use enums instead of getcollection before read / writes
func (db *Database) getCollection(name string) *mongo.Collection {
	return db.client.Database(db.config.DBSchema).Collection(name)
}

func (db *Database) Write(
	collectionName  string,
	object      		any,
) (*mongo.InsertOneResult, error) {
	collection := db.getCollection(collectionName)
	result, err := collection.InsertOne(context.TODO(), object)
	return result, err
}

func (db *Database) FindOne(
	collectionName  string,
	query       		*bson.D,
) *mongo.SingleResult {
	collection := db.getCollection(collectionName)
	return collection.FindOne(context.TODO(), query)
}

func (db *Database) UpdateOne(
	collectionName 	string,
	filter 					*bson.D,
	query 					*bson.D,
) error {
	collection := db.getCollection(collectionName)
	_, err := collection.UpdateOne(context.TODO(), filter, query)
	return err
}

func (db *Database) ClearCollection(collectionName string) (*mongo.DeleteResult, error) {
	collection := db.getCollection(collectionName)
	result, err := collection.DeleteMany(context.TODO(), bson.D {})
	return result, err
}

func (db *Database) GetAllNodes() ([]*node.Node, error) {
	collection := db.getCollection(NODES_TABLE)
	cursor, err := collection.Find(context.TODO(), bson.D {})
	if err != nil {
		return nil, err
	}
	var nodes []*node.Node
	if err = cursor.All(context.TODO(), &nodes); err != nil {
		return nil, err
	}
	return nodes, nil
}

func (db *Database) SetNodeLastHB(node *node.Node) error {
	return db.UpdateOne(
		NODES_TABLE,
		&bson.D{{Key: "nodeId", Value: node.ID}},
		&bson.D{{Key: "$set", Value: bson.D{{Key: "lastHB", Value: time.Now().Unix()}}}},
	)
}

func (db *Database) SetNodeLastCommandAndResponse(node *node.Node, command string, response string) error {
	return db.UpdateOne(
		NODES_TABLE,
		&bson.D{{Key: "nodeId", Value: node.ID}},
		&bson.D{{Key: "$set", Value: bson.D{{Key: "lastCommand", Value: command}, {Key: "lastResponse", Value: response}}}},
	)
}