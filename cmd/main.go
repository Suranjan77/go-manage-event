package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	props "github.com/Suranjan77/go-manage-event/pkg/common/config"
	"github.com/Suranjan77/go-manage-event/pkg/common/db"
	"github.com/Suranjan77/go-manage-event/pkg/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

	router.Use(cors.Default())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong", "currentTime": time.Now().Format("2006-01-02T15:04:05+07:00")})
	})

	db := setupDb()
	routes.RegisterRoutes(router, db)

	return router
}

func setupDb() *gorm.DB {
	dbProps := props.P.DataSource

	dsn := fmt.Sprintf(
		"%v:%v@tcp(%v:%v)/%v",
		dbProps.UserName,
		dbProps.Password,
		dbProps.Host,
		dbProps.Port,
		dbProps.DbName,
	)

	return db.Init(dsn)
}

func main() {
	router := setUpRouter()

	log.Printf("Server started on port: %v \n", props.P.Server.Port)

	router.Run(":" + props.P.Server.Port)
}
