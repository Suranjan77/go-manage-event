package main

import (
	"net/http"
	"strconv"
	"time"

	config "github.com/Suranjan77/go-manage-event/pkg/config"
	"github.com/sirupsen/logrus"
	"github.com/toorop/gin-logrus"

	"github.com/Suranjan77/go-manage-event/pkg/db"
	"github.com/Suranjan77/go-manage-event/pkg/middlewares"
	"github.com/Suranjan77/go-manage-event/pkg/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var log *logrus.Logger

func setUpRouter() *gin.Engine {
	router := gin.New()

	router.Use(ginlogrus.Logger(log))

	router.Use(gin.Recovery())

	router.Use(cors.Default())

	router.Use(middlewares.JSONMiddleware())

	router.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"message": "pong", "currentTime": time.Now().Format("2006-01-02T15:04:05+07:00")})
	})

	routes.RegisterRoutes(router, db.GetDB())

	return router
}

func main() {

	config.Load()

	log = config.Logger()

	db.SetupDB()

	router := setUpRouter()

	log.Info("Server started on port: %v", config.Port())

	router.Run(":" + strconv.Itoa(config.Port()))
}
