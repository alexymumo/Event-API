package main

import (
	"events/internal/routes"
	"events/pkg/database"
	"events/pkg/util"

	"github.com/gin-gonic/gin"
)

func init() {
	database.Connect()
	util.LoadEnv()
}
func main() {
	router := gin.New()
	router.Use(gin.Logger())
	routes.EventRoutes(router)
	routes.UserRoutes(router)
	router.Run("localhost:8080")
}
