package main

import (
	node "einsof/sauron/pkg"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type API struct {
	Engine *gin.Engine
	db     *Database
}

func (api *API) Init(db *Database) {
	api.db      = db
	api.Engine  = gin.Default()

	api.Engine.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:3000"},
		AllowMethods:     []string{"OPTIONS", "PUT", "PATCH", "GET", "POST"},
		AllowHeaders:     []string{"Origin", "Access-Control-Allow-Origin"},
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

func (api *API) SetRoutes() {
	api.Engine.GET("/", api.healthCheckFunc)
	api.Engine.GET("/client/register", api.registerNewNodeFunc)
	api.Engine.GET("/nodes", api.getAllNodesFunc)
}