package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log/slog"

	node "einsof/sauron/pkg"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/packets"
)

type MQTTHookOptions struct {
	Server   *mqtt.Server
	Database *Database
	Log      *slog.Logger
}

type MQTTHook struct {
	mqtt.HookBase
	config *MQTTHookOptions
	log    *slog.Logger
	db     *Database
}

func (hook *MQTTHook) ID() string {
	return "MQTTHook"
}

func (hook *MQTTHook) Provides(b byte) bool {
	return bytes.Contains(
		[]byte{
			mqtt.OnConnect,
			mqtt.OnDisconnect,
			mqtt.OnSubscribed,
			mqtt.OnUnsubscribed,
			mqtt.OnPublish,
			mqtt.OnPublished,
		},
		[]byte{b},
	)
}

func (hook *MQTTHook) Publish(topic string, message []byte) {
	hook.config.Server.Publish(topic, message, false, 1)
}

func (hook *MQTTHook) Init(config any) error {
	if _, ok := config.(*MQTTHookOptions); !ok && config != nil {
		return mqtt.ErrInvalidConfigType
	}

	hook.config = config.(*MQTTHookOptions)
	hook.log    = hook.config.Log
	hook.db     = hook.config.Database

	if hook.config.Server == nil {
		return mqtt.ErrInvalidConfigType
	}

	return nil
}

func (hook *MQTTHook) OnConnect(
	client *mqtt.Client,
	packet packets.Packet,
) error {
	hook.log.Info("client connected", "client", client.ID)

	nodeObj := GetNodeFromDB(client.ID, hook.db)

	if nodeObj != nil {
		hook.log.Info(
			"client from db",
			"client", nodeObj,
		)
		// SetNodeOnlineStatus(nodeObj.ID, true, hook.db)
	} else {
		hook.log.Warn("could not find node with ID in database", "client", client.ID)
		client.Stop(nil)
	}

	// newNode := new(node.Node)
	// newNode.SetId("testid")
	// newNode.SetIPAddress(client.Net.Conn.RemoteAddr().String())
	// hook.config.CoreNode.AddNode(newNode)

	return nil
}

func (hook *MQTTHook) OnDisconnect(
	client  *mqtt.Client,
	err     error,
	expire  bool,
) {
	nodeObj := GetNodeFromDB(client.ID, hook.db)

	if err != nil {
		hook.log.Info(
			"client disconnected",
			"client", client.ID,
			"expire", expire,
			"error",  err,
		)
	} else {
		hook.log.Info(
			"client disconnected",
			"client", client.ID,
			"expire", expire,
		)
	}

	if nodeObj != nil {
		hook.log.Info(
			"client from db",
			"client", nodeObj,
		)
		SetNodeOnlineStatus(nodeObj.ID, false, hook.db)
	} else {
		hook.log.Warn("could not find node with ID in database", "client", client.ID)
	}
}

func (hook *MQTTHook) OnSubscribed(
	client      *mqtt.Client,
	packet      packets.Packet,
	reasonCodes []byte,
) {
	hook.log.Info(
		fmt.Sprintf("client subscribed qos=%v", reasonCodes),
		"client",   client.ID,
		"filters",  packet.Filters,
	)

	topic := GetNodeTopic(client.ID, hook.db)

	if topic != packet.Filters[0].Filter {
		hook.log.Warn(
			fmt.Sprintf(
				"node %s subscribed on wrong topic %s, closing connection",
				client.ID, packet.Filters[0].Filter,
			),
		)
		client.Stop(nil)
	} else {
		msg := node.Message{}
		msg.Type = "netstat"
		msg.Data = ""
		jsonBytes, _ := json.Marshal(msg)
		hook.Publish(topic, jsonBytes)
	}
}

func (hook *MQTTHook) OnUnsubscribed(
	client *mqtt.Client,
	packet packets.Packet,
) {
	hook.log.Info(
		"client unsubscribed",
		"client",   client.ID,
		"filters",  packet.Filters,
	)
}

func (hook *MQTTHook) OnPublish(
	client *mqtt.Client,
	packet packets.Packet,
) (packets.Packet, error) {
	hook.log.Info(
		"received from client",
		"client",   client.ID,
		"payload",  string(packet.Payload),
	)

	packetx := packet

	fmt.Println(string(packet.Payload))

	return packetx, nil
}

func (hook *MQTTHook) OnPublished(
	client *mqtt.Client,
	packet packets.Packet,
) {
	hook.log.Info(
		"published to client",
		"client",   client.ID,
		"payload",  string(packet.Payload),
	)
}
