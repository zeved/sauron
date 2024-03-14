package main

import (
	node "einsof/sauron/pkg"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	mqtt "github.com/mochi-mqtt/server/v2"
)

type API struct {
	Engine *gin.Engine
	db     *Database
	mqtt   *mqtt.Server
}

type APIMQTTCmdReq struct {
	Command string `json:"command"`
	NodeId  string `json:"nodeId"`
}

func (api *API) Init(db *Database, mqtt *mqtt.Server) {
	api.db      = db
	api.Engine  = gin.Default()
	api.mqtt 		= mqtt

	api.Engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"OPTIONS", "PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin", "Content-Type"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		// MaxAge: 12 * .Hour,
}))
}

func (api *API) healthCheckFunc(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})
}

func (api *API) registerNewNodeFunc(ctx *gin.Context) {
	var nodeObj node.Node
	nodeId, err := uuid.NewRandom()
	if err != nil {
		return
	}

	nodeObj.ID = nodeId.String()
	nodeObj.Topic = fmt.Sprintf("node/%s", nodeId.String())
	

	api.db.Write("nodes", nodeObj)
	// err := ctx.BindJSON(&d)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// fmt.Println(d.ClientId)
	ctx.JSON(http.StatusCreated, nodeObj)
}


func (api *API) getAllNodesFunc(ctx *gin.Context) {
	nodes, err := api.db.GetAllNodes()

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusOK, nodes)
}

func (api *API) getNodefunc(ctx *gin.Context) {

	nodeId := ctx.Param("id")
	nodeObj := api.db.GetNode(nodeId)
	ctx.JSON(http.StatusOK, nodeObj)
}

func (api *API) sendNodeCmdFunc(ctx *gin.Context) {
	mqttCmd := new(APIMQTTCmdReq)
	err := ctx.BindJSON(mqttCmd)

	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}

	fmt.Printf("mqttCmd: %+v\n", mqttCmd)

	msg := node.Message{
		Type: mqttCmd.Command,
	}

	json, _ := json.Marshal(msg)

	api.mqtt.Publish(fmt.Sprintf("node/%s", mqttCmd.NodeId), json, false, 1)

	ctx.JSON(http.StatusOK, gin.H{
		"message": "OK",
	})	
}

func (api *API) SetRoutes() {
	api.Engine.GET("/", api.healthCheckFunc)
	api.Engine.GET("/client/register", api.registerNewNodeFunc)
	api.Engine.GET("/nodes", api.getAllNodesFunc)
	api.Engine.GET("/nodes/:id", api.getNodefunc)
	api.Engine.POST("/nodes/cmd", api.sendNodeCmdFunc)
}