package main

import (
	"fmt"
	"net/http"
	"time"

	props "github.com/Suranjan77/go-manage-event/utils"
	"github.com/gin-gonic/gin"
)

func setUpRouter() *gin.Engine {
	router := gin.New()

	/* 	// Write logs to file and print in the standard output
	   	logFile, err := os.Create("/var/log/go-manage-event-server.log")

	   	if err != nil {
	   		fmt.Printf("Failed to create logs file %v", err.Error())
	   	}

	   	gin.DisableConsoleColor()
	   	router.Use(gin.LoggerWithWriter(io.MultiWriter(logFile, os.Stdout))) */

	router.Use(gin.Logger())

	router.Use(gin.Recovery())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong", "currentTime": time.Now().Format("2006-01-02T15:04:05+07:00")})
	})

	return router
}

func main() {
	router := setUpRouter()

	fmt.Printf("Server started on port: %v \n", props.P.Server.Port)

	router.Run(":" + props.P.Server.Port)

}
