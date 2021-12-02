package main

import (
	"github.com/ahmedsharyo/SystemContorellerService/router"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.New()

	router.Routering(server)

	port := "8080"
	server.Run(":" + port)
}
