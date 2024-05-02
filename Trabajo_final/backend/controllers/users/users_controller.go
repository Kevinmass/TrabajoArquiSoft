package users

import (
	userService "backend/services/users" //se le asigna un alias porque tiene el mismo nombre que el domain

	"net/http"

	"gitgub.com/gin-gonic/gin"

	userDomain "backend/domain/users"
)

func Login(c *gin.Context) {

	var userData userDomain.UserData
	c.BindJSON(&userData)
	loginResponse := userService.Login(userData.User, userData.Password)
	c.JSON(http.StatusOK, loginResponse)
}
