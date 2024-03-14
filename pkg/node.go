package node

type Node struct {
	ID          string `bson:"nodeId"`
	IPAddress   string `bson:"ipAddress"`
	Topic       string `bson:"topic"`
	Online      bool   `bson:"online"`
	JoinedOn    int64  `bson:"joinedOn"`
	LastHB      int64  `bson:"lastHB"`
	LastCommand string `bson:"lastCommand"`
}

func (node *Node) SetIPAddress(ipAddress string) {
	node.IPAddress = ipAddress
}

func (node *Node) SetTopic(topic string) {
	node.Topic = topic
}