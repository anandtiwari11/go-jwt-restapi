package main

import (
	"log"
	"github.com/gin-gonic/gin"
	"gojwtproject/config"
	"gojwtproject/routes"
)

func main() {
	config.ConnectDB()
	router := gin.Default()
	routes.AuthRoutes(router)
	log.Fatal(router.Run(":8080"))
}