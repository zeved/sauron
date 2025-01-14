package main

import (
	node "einsof/sauron/pkg"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

func GetNodeFromDB(nodeId string, db *Database) *node.Node {
	nodeFromDB := db.FindOne(NODES_TABLE, &bson.D {
		{ Key: "nodeId",  Value: nodeId },
		// { Key: "active",    Value: true },
	})

	if nodeFromDB.Err() != mongo.ErrNoDocuments {
		var nodeObj node.Node
		nodeFromDB.Decode(&nodeObj)
		return &nodeObj
	} else {
		return nil
	}
}

func SetNodeOnlineStatus(nodeId string, online bool, db *Database) error {
	filter := bson.D {{ Key: "nodeId", Value: nodeId }}
	query := bson.D {{ Key: "$set", Value: bson.D {{ Key: "online", Value: online }}}}	

	return db.UpdateOne(NODES_TABLE, &filter, &query)
}

func SetNodeTopic(nodeId string, topic string, db *Database) error {
	filter := bson.D {{ Key: "nodeId", Value: nodeId }}
	query := bson.D {{ Key: "$set", Value: bson.D {{ Key: "topic", Value: topic }}}}	

	return db.UpdateOne(NODES_TABLE, &filter, &query)
}

func GetNodeTopic(nodeId string, db *Database) string {
	nodeFromDB := GetNodeFromDB(nodeId, db)
	return nodeFromDB.Topic
}

func SendMessageToNode(nodeId string, command string, db *Database) {
	nodeFromDB := GetNodeFromDB(nodeId, db)

	if (nodeFromDB != nil) {
		
	} else {
		return
	}
}

func SaveNodeHostInfo(nodeId string, hostInfo string, db *Database) {
	filter := bson.D {{ Key: "nodeId", Value: nodeId }}
	query := bson.D {{ Key: "$set", Value: bson.D {{ Key: "hostInfo", Value: hostInfo }}}}

	db.UpdateOne(NODES_TABLE, &filter, &query)
}

func SaveNodeCPUInfo(nodeId string, cpuInfo string, db *Database) {
	filter := bson.D {{ Key: "nodeId", Value: nodeId }}
	query := bson.D {{ Key: "$set", Value: bson.D {{ Key: "cpuInfo", Value: cpuInfo }}}}

	db.UpdateOne(NODES_TABLE, &filter, &query)
}

func SaveNodeMemInfo(nodeId string, memInfo string, db *Database) {
	filter := bson.D {{ Key: "nodeId", Value: nodeId }}
	query := bson.D {{ Key: "$set", Value: bson.D {{ Key: "memInfo", Value: memInfo }}}}

	db.UpdateOne(NODES_TABLE, &filter, &query)
}

func SaveNodeProcListInfo(nodeId string, processList string, db *Database) {
	filter := bson.D {{ Key: "nodeId", Value: nodeId }}
	query := bson.D {{ Key: "$set", Value: bson.D {{ Key: "processList", Value: processList }}}}

	db.UpdateOne(NODES_TABLE, &filter, &query)
}

func SaveNodeNetConnListInfo(nodeId string, netConnList string, db *Database) {
	filter := bson.D {{ Key: "nodeId", Value: nodeId }}
	query := bson.D {{ Key: "$set", Value: bson.D {{ Key: "netConnList", Value: netConnList }}}}

	db.UpdateOne(NODES_TABLE, &filter, &query)
}