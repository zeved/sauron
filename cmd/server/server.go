package main

import (
	node "einsof/sauron/pkg"
	"fmt"
	"log/slog"
	"os"

	"github.com/spf13/viper"
)

type Config struct {
	HTTPPort string
	MQTTPort string
	DBHost   		string
	DBPort   		string
	DBSchema 		string
	DBUsername 	string
	DBPassword	string
}

type Server struct {
	config *Config
	Log *slog.Logger
	nodes []*node.Node
	db *Database
}

func (srv *Server) Init() error {
	srv.Log = slog.New(slog.NewTextHandler(os.Stdout, nil))
	viper.SetConfigFile("sauron.json")
	viper.SetConfigType("json")
	viper.AddConfigPath(".")

	err := viper.ReadInConfig()

	if err != nil {
		panic(fmt.Errorf("failed reading config file %w", err))
	}

	srv.config = new(Config)

	srv.config.HTTPPort = viper.GetString("http.port")
	srv.config.MQTTPort = viper.GetString("mqtt.port")
	srv.config.DBHost    	= viper.GetString("database.host")
	srv.config.DBSchema  	= viper.GetString("database.schema")
	srv.config.DBPort    	= viper.GetString("database.port")
	srv.config.DBUsername 	= viper.GetString("database.username")
	srv.config.DBPassword  = viper.GetString("database.password")

	srv.db = new(Database)
	srv.db.Init(srv.config)

	return nil
}

func (srv *Server) Config() *Config {
	return srv.config
}

func (srv *Server) AddNode(node *node.Node) {
	srv.nodes = append(srv.nodes, node)
}