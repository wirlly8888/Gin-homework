package main

import (
	"github.com/gin-gonic/gin"

	database "homework/data_base"
	"homework/handler"
)

func main() {
	Db := database.NewTempDataBase()
	server := gin.Default()
	server = handler.SetupRoute(server, Db)
	server.Run(":8080")
}
