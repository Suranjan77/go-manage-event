package routes

import (
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"

	"github.com/Suranjan77/go-manage-event/pkg/routes/handler"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	h := handler.Handler{
		DB: db,
	}

	publicUsrGrp := r.Group("/public/users")
	addPublicUserRoutes(publicUsrGrp, h)

}

func addPublicUserRoutes(route *gin.RouterGroup, h handler.Handler) {
	route.POST("/signup", h.AddUser)
}
