package main

import (
	"Server/handler"
	"github.com/gin-gonic/gin"
)

var bioHandler *handler.Handler

func main() {
	router := gin.Default()
	router.GET("/", bioHandler.BioHandler)

	router.Run()
}
