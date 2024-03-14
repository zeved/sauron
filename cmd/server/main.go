package main

import (
	"log"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	mqtt "github.com/mochi-mqtt/server/v2"
	"github.com/mochi-mqtt/server/v2/hooks/auth"
	"github.com/mochi-mqtt/server/v2/listeners"
)

func main() {
	server := new(Server)
	server.Init()

	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		<-sigs
		done <- true
	}()

	level := new(slog.LevelVar)
	mqttServer := mqtt.New(&mqtt.Options{
		InlineClient: true,
	})
	mqttServer.Log = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
		Level: level,
	}))
	level.Set(slog.LevelWarn)

	_ = mqttServer.AddHook(new(auth.AllowHook), nil)

	tcp := listeners.NewTCP("t1", ":1883", nil)
	err := mqttServer.AddListener(tcp)
	if err != nil {
		log.Fatal(err)
	}

	err = mqttServer.AddHook(new(MQTTHook), &MQTTHookOptions{
		Server:   mqttServer,
		Database: server.db,
		Log:      server.Log,
	})

	if err != nil {
		log.Fatal(err)
	}

	api := new(API)
	api.Init(server.db)
	api.SetRoutes()

	go func() {
		err = mqttServer.Serve()
		if err != nil {
			log.Fatal(err)
		}
		api.Engine.Run()
	}()

	<-done

	err = mqttServer.Close()
	if err != nil {
		log.Fatal(err)
	}
}