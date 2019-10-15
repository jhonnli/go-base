package users

import (
	"github.com/gin-gonic/gin"
	"github.com/jhonnli/go-base/controller/common"
	userService "github.com/jhonnli/go-base/service/users"
	"net/http"
)

func InitUsers(engin *gin.Engine) *gin.RouterGroup {
	userApi := engin.Group("/v1/users")
	common.AddFilter(userApi)
	userApi.GET("", getUsers)
	return userApi
}

func getUsers(context *gin.Context) {
	users := userService.GetUsers()
	context.JSON(http.StatusOK, common.GenSuccessResult(users))
}
