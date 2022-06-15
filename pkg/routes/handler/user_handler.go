package handler

import (
	"net/http"
	"time"

	"github.com/Suranjan77/go-manage-event/pkg/models"
	"github.com/gin-gonic/gin"
)

func (h *Handler) AddUser(c *gin.Context) {
	signUpRequest := models.SignUpRequest{}

	if err := c.BindJSON(&signUpRequest); err != nil {
		error := models.ErrorResponse{
			Error: []models.Error{
				{Msg: err.Error()},
			},
			TimeStamp: time.Now().UnixMilli(),
		}
		c.AbortWithStatusJSON(http.StatusBadRequest, &error)
		return
	}

	user := signUpRequest.ToUser()

	if res := h.DB.Create(&user); res.Error != nil {
		c.AbortWithStatusJSON(
			http.StatusNotFound,
			models.ErrorResponse{
				Error: []models.Error{
					{Msg: res.Error.Error()},
				},
				TimeStamp: time.Now().UnixMilli(),
			},
		)
		return
	}

	c.JSON(http.StatusCreated, models.UserResponseFromUser(user))
}
