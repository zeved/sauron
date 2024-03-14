package main

import (
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"time"

	node "einsof/sauron/pkg"

	mqtt "github.com/eclipse/paho.mqtt.golang"
	"github.com/spf13/viper"
)

var message = make(chan []string)

var f mqtt.MessageHandler = func(client mqtt.Client, msg mqtt.Message) {
	message <- []string{msg.Topic(), string(msg.Payload())}
}

func handleMessage(message string, client mqtt.Client, clientId string) {
	msg := node.Message{}
	json.Unmarshal([]byte(message), &msg)

	msgReply := node.Message{
		Timestamp: time.Now().Unix(),
	}

	switch msg.Type {
	case "command":
		fmt.Printf("command: %s\n", msg.Data)
		out, _ := exec.Command("bash", "-c", msg.Data).Output()
		client.Publish(fmt.Sprintf("node/%s", clientId), 0, false, string(out))

	case "ps":
		msgData, err := node.GetProcessList()

		if err != nil {
			fmt.Println(err)
		} else {
			msgReply.Data = msgData
		}

	case "cpu":
		msgReply.Data = node.GetCPUInfo()

	case "mem":
		mem, _ := node.GetMemoryInfo()
		msgReply.Data = mem

	case "netstat":
		netstat, _ := node.GetNetstatInfo()
		msgReply.Data = netstat
		
	case "hostinfo":
		hostInfo := node.GetHostInfo()
		msgReply.Data = hostInfo

	case "leave":
		client.Disconnect(0)
	case "reply":
		break
	}

	if (msgReply.Data == "") {
		return
	}

	msgReply.Type = fmt.Sprintf("%s.reply", msg.Type)
	jsonBytes, _ := json.Marshal(msgReply)
	client.Publish(fmt.Sprintf("node/%s", clientId), 0, false, jsonBytes)
}

func main() {
	viper.SetConfigFile("sauron-client.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("failed reading config file %w", err))
	}

	config := new(Config)
	config.ID = viper.GetString("id")
	config.serverIP = viper.GetString("server.ip")
	config.serverPort = viper.GetString("server.port")


	fmt.Printf("ID: %s\n", config.ID)
	fmt.Printf("Server IP: %s\n", config.serverIP)
	fmt.Printf("Server Port: %s\n", config.serverPort)

	opts := mqtt.NewClientOptions().AddBroker(
		fmt.Sprintf("tcp://%s:%s", config.serverIP, config.serverPort)).
		SetClientID(config.ID)

	opts.SetKeepAlive(10 * time.Second)
	opts.SetDefaultPublishHandler(f)
	opts.SetPingTimeout(1 * time.Second)

	client := mqtt.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		panic(token.Error())
	}

	fmt.Printf("Subscribing to topic %s\n", "sauron")
	topic := fmt.Sprintf("node/%s", config.ID)
	if token := client.Subscribe(topic, 0, f); token.Wait() && token.Error() != nil {
		fmt.Println(token.Error())
		os.Exit(1)
	}

	for {
		incoming := <- message
		handleMessage(incoming[1], client, config.ID)
	}

	// client.Disconnect(0)
}