package node

type Node struct {
	ID           string `bson:"nodeId"`
	IPAddress    string `bson:"ipAddress"`
	Topic        string `bson:"topic"`
	Online       bool   `bson:"online"`
	JoinedOn     int64  `bson:"joinedOn"`
	LastHB       int64  `bson:"lastHB"`
	LastCommand  string `bson:"lastCommand"`
	LastResponse string `bson:"lastResponse"`
	HostInfo     string `bson:"hostInfo"`
	CPUInfo      string `bson:"cpuInfo"`
	MemInfo      string `bson:"memInfo"`
	ProcessList  string `bson:"processList"`
	NetConnList  string `bson:"netConnList"`
}

func (node *Node) SetIPAddress(ipAddress string) {
	node.IPAddress = ipAddress
}

func (node *Node) SetTopic(topic string) {
	node.Topic = topic
}
