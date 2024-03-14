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

	switch msg.Type {
	case "command":
		fmt.Printf("command: %s\n", msg.Data)
		out, _ := exec.Command("bash", "-c", msg.Data).Output()
		client.Publish(fmt.Sprintf("node/%s", clientId), 0, false, string(out))

	case "ps":
		processes, err := node.GetProcessList()

		if err != nil {
			fmt.Println(err)
		}

		output, _ := json.Marshal(processes)
		client.Publish(fmt.Sprintf("node/%s", clientId), 0, false, output)

	case "cpu":
		cpuInfo := node.GetCPUInfo()
		fmt.Println(cpuInfo)
		client.Publish(fmt.Sprintf("node/%s", clientId), 0, false, cpuInfo)

	case "mem":
		mem, _ := node.GetMemoryInfo()
		fmt.Println(mem)
		client.Publish(fmt.Sprintf("node/%s", clientId), 0, false, mem)

	case "netstat":
		netstat, _ := node.GetNetstatInfo()
		client.Publish(fmt.Sprintf("node/%s", clientId), 0, false, netstat)
	
	case "leave":
		client.Disconnect(0)
	case "reply":
		break
	}
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

	opts.SetKeepAlive(2 * time.Second)
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
		// fmt.Printf("Received message on topic %s: %s\n", incoming[0], incoming[1])
		handleMessage(incoming[1], client, config.ID)
	}

	// client.Disconnect(0)
}