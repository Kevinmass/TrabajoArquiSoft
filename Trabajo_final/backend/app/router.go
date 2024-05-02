package app

import (
	"backend/controllers/users"

	"github.com/gin-gonic-gin"
)

func MapRoutes(engines *gin.Engine) {
	engine.POST("/users/login", users.Login)
}
