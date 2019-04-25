package main

import (
	"github.com/gin-gonic/gin"

	"github.com/RaghavHarithas/2019-04-friendbook-profile/database/connection"
	"github.com/RaghavHarithas/2019-04-friendbook-profile/database/schema"
	"github.com/RaghavHarithas/2019-04-friendbook-profile/handlers"
)

func main() {
	schema.Init()
	connection.ConnectPool()
	defer connection.DisconnectPool()

	r := gin.Default()
	r.GET("/profile", handlers.GetProfile)
	r.PUT("/profile", handlers.CreateOrUpdateProfile)
	r.DELETE("/profile", handlers.DeleteProfile)
	r.Run() // listen and serve on 0.0.0.0:8080
}
