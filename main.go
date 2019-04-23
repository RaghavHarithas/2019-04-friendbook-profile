package main

import (
	"github.com/gin-gonic/gin"

	"github.com/rajch/2019-04-friendbook-profile/database/connection"
	"github.com/rajch/2019-04-friendbook-profile/database/schema"
	"github.com/rajch/2019-04-friendbook-profile/handlers"
)

func main() {
	schema.Init()
	connection.ConnectPool()
	defer connection.DisconnectPool()

	r := gin.Default()
	r.GET("/profiles", handlers.GetProfile)
	r.PUT("/profiles", handlers.CreateOrUpdateProfile)
	r.Run() // listen and serve on 0.0.0.0:8080
}
